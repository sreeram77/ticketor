package main

import (
	"log/slog"
	"net"
	"os"

	"ticketor/cmd/server/handlers"
	protogen "ticketor/protogen/proto"
	"ticketor/store"

	"google.golang.org/grpc"
)

func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug)
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	})))

	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		slog.Error("failed to listen", "err", err)
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
		slog.Error("failed to serve", "err", err)
	}
}
