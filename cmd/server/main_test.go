package main

import (
	"context"
	"log"
	"net"
	"os"
	"testing"

	"ticketor/cmd/server/handlers"
	protogen "ticketor/protogen/proto"
	"ticketor/store"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

var grpcServer *grpc.Server
var serverAddress string

func TestMain(m *testing.M) {
	lis, err := net.Listen("tcp", ":0") // Random available port
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	serverAddress = lis.Addr().String()

	grpcServer = grpc.NewServer()

	userStore := store.NewUsers()

	ticketorHandler := handlers.NewTicketor(userStore, store.NewSections(), store.NewTickets())
	userHandler := handlers.NewUser(userStore)

	protogen.RegisterUserServiceServer(grpcServer, userHandler)
	protogen.RegisterTicketorServer(grpcServer, ticketorHandler)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	// Run the tests
	code := m.Run()

	// Shutdown server
	grpcServer.GracefulStop()
	os.Exit(code)
}

func userServiceClient(t *testing.T) protogen.UserServiceClient {
	conn, err := grpc.NewClient(serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		t.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	t.Cleanup(func() { conn.Close() })

	return protogen.NewUserServiceClient(conn)
}

func ticketorServiceClient(t *testing.T) protogen.TicketorClient {
	conn, err := grpc.NewClient(serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		t.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	t.Cleanup(func() { conn.Close() })

	return protogen.NewTicketorClient(conn)
}

func TestServer(t *testing.T) {
	userClient := userServiceClient(t)
	ticketorClient := ticketorServiceClient(t)

	type testCase struct {
		name        string
		user        *protogen.UserRequest
		ticket      *protogen.TicketRequest
		created     *protogen.UserResponse
		ticketReply *protogen.TicketReply
		err         error
	}

	testCases := []testCase{
		{
			name: "John Doe",
			user: &protogen.UserRequest{
				FirstName: "John",
				LastName:  "Doe",
				Email:     "john@doe.com",
			},
			ticket: &protogen.TicketRequest{
				From: "France",
				To:   "London",
			},
		},
		{
			name: "Jane Doe",
			user: &protogen.UserRequest{
				FirstName: "Jane",
				LastName:  "Doe",
				Email:     "jane@doe.com",
			},
			ticket: &protogen.TicketRequest{
				From: "France",
				To:   "London",
			},
		},
	}

	for _, tc := range testCases {
		// Create user.
		created, err := userClient.Create(context.Background(), tc.user)
		if err != nil {
			t.Fatal(err)
		}

		if created.Id == "" {
			t.Errorf("expected ID to be set")
		}

		if created.Email != tc.user.Email {
			t.Errorf("expected %v, got %v", tc.user, created)
		}

		tc.created = created
		tc.ticket.UserId = created.Id

		// Purchase ticket.
		ticket, err := ticketorClient.PurchaseTicket(context.Background(), tc.ticket)
		if err != nil {
			t.Fatal(err)
		}

		if ticket.Id == "" {
			t.Errorf("expected ID to be set")
		}

		if ticket.From != tc.ticket.From {
			t.Errorf("expected %v, got %v", tc.ticket, ticket)
		}

		if ticket.To != tc.ticket.To {
			t.Errorf("expected %v, got %v", tc.ticket, ticket)
		}

		tc.ticketReply = ticket

		// Get ticket.
		getTicket, err := ticketorClient.GetTicket(context.Background(), &protogen.TicketIDRequest{Id: ticket.Id})
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		if getTicket.Id != ticket.Id {
			t.Errorf("expected %v, got %v", ticket, getTicket)
		}

		if getTicket.User == nil {
			t.Errorf("expected user to be set")
		}

		if getTicket.User.Id != created.GetId() {
			t.Errorf("expected %v, got %v", created, getTicket.GetUser())
		}

		// Modify ticket.
		modifiedTicket, err := ticketorClient.ModifyTicket(context.Background(), &protogen.TicketRequest{
			UserId: created.GetId(),
			From:   tc.ticket.From,
			To:     tc.ticket.To,
			Id:     tc.ticketReply.Id,
		})
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		if tc.ticketReply.GetSection() == modifiedTicket.GetSection() && tc.ticketReply.GetNumber() == modifiedTicket.GetNumber() {
			t.Errorf("expected ticket number to be modified")
		}

		// Remove ticket.
		_, err = ticketorClient.RemoveTicket(context.Background(), &protogen.TicketIDRequest{Id: tc.ticketReply.Id})
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		_, err = ticketorClient.GetTicket(context.Background(), &protogen.TicketIDRequest{Id: ticket.Id})
		if err == nil {
			t.Errorf("expected error, got nil")
		}

		st, ok := status.FromError(err)
		if !ok {
			t.Errorf("expected status, got nil")
		}

		if st.Code() != codes.NotFound {
			t.Errorf("expected %v, got %v", codes.NotFound, st.Code())
		}
	}
}
