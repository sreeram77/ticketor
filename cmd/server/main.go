package main

import (
	"log/slog"
	"net"

	"ticketor/cmd/server/handlers"
	protogen "ticketor/protogen/proto"
	"ticketor/store"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		slog.Error("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	userStore := store.NewUsers()

	ticketorHandler := handlers.NewTicketor(userStore, store.NewSections(), store.NewTickets())
	userHandler := handlers.NewUser(userStore)

	protogen.RegisterUserServiceServer(grpcServer, userHandler)
	protogen.RegisterTicketorServer(grpcServer, ticketorHandler)

	err = grpcServer.Serve(lis)
	if err != nil {
		slog.Error("failed to serve: %v", err)
	}
}
