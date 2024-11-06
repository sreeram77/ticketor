package main

import (
	"log/slog"
	"net"

	"ticketor/handlers"
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

	ticketorHandler := handlers.NewTicketor(store.NewUsers(), store.NewSections(), store.NewTickets())

	protogen.RegisterTicketorServer(grpcServer, ticketorHandler)

	err = grpcServer.Serve(lis)
	if err != nil {
		slog.Error("failed to serve: %v", err)
	}
}
