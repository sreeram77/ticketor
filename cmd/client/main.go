package main

import (
	"context"
	"log/slog"
	"os"

	protogen "ticketor/protogen/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug)
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	})))

	conn, err := grpc.NewClient("localhost:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		slog.Error("failed to dial:", "error", err)
	}

	userClient := protogen.NewUserServiceClient(conn)
	ticketorClient := protogen.NewTicketorClient(conn)

	createdUser, err := userClient.Create(context.Background(), &protogen.UserRequest{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@doe.com",
	})
	if err != nil {
		slog.Error("failed to create user:", "error", err)
		return
	}

	ticket, err := ticketorClient.PurchaseTicket(context.Background(), &protogen.TicketRequest{
		UserId: createdUser.Id,
		From:   "France",
		To:     "London",
	})
	if err != nil {
		slog.Error("failed to purchase ticket:", "error", err)
		return
	}

	slog.Info("purchase ticket", "ticket", ticket)

	tickets, err := ticketorClient.GetTickets(context.Background(), &protogen.SectionIDRequest{
		Id: ticket.Section,
	})
	if err != nil {
		slog.Error("failed to get tickets:", "error", err)
		return
	}

	slog.Info("tickets: %v", tickets)

	_, err = ticketorClient.RemoveTicket(context.Background(), &protogen.TicketIDRequest{
		Id: ticket.Id,
	})
	if err != nil {
		slog.Error("failed to remove ticket:", "error", err)
		return
	}

	createdUser2, err := userClient.Create(context.Background(), &protogen.UserRequest{
		FirstName: "Johnny",
		LastName:  "Doe",
		Email:     "johnny@doe.com",
	})
	if err != nil {
		slog.Error("failed to create user:", "error", err)
		return
	}

	createdUser3, err := userClient.Create(context.Background(), &protogen.UserRequest{
		FirstName: "Jane",
		LastName:  "Doe",
		Email:     "jane@doe.com",
	})
	if err != nil {
		slog.Error("failed to create user:", "error", err)
		return
	}

	ticket2, err := ticketorClient.PurchaseTicket(context.Background(), &protogen.TicketRequest{
		UserId: createdUser2.Id,
		From:   "France",
		To:     "London",
	})
	if err != nil {
		slog.Error("failed to purchase ticket:", "error", err)
		return
	}

	_, err = ticketorClient.PurchaseTicket(context.Background(), &protogen.TicketRequest{
		UserId: createdUser3.Id,
		From:   "France",
		To:     "London",
	})
	if err != nil {
		slog.Error("failed to purchase ticket:", "error", err)
		return
	}

	modifyTicket2, err := ticketorClient.ModifyTicket(context.Background(), &protogen.TicketRequest{
		Id: ticket2.Id,
	})
	if err != nil {
		slog.Error("failed to modify ticket:", "error", err)
		return
	}

	slog.Info("modify ticket:", "ticket", modifyTicket2)

	tickets, err = ticketorClient.GetTickets(context.Background(), &protogen.SectionIDRequest{
		Id: ticket.Section,
	})
	if err != nil {
		slog.Error("failed to get tickets:", "error", err)
		return
	}

	slog.Info("tickets by section:", "len", len(tickets.Tickets), "tickets", tickets.Tickets)
}
