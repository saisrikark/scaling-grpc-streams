package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"scaling-grpc-streams/pkg/application"

	"google.golang.org/grpc"
)

type Server struct {
	application.UnimplementedGreeterServer
	grpcServer *grpc.Server
	port       int
}

func NewServer(port int) (*Server, error) {
	server := &Server{port: port}

	grpcServer := grpc.NewServer()
	application.RegisterGreeterServer(grpcServer, server)

	server.grpcServer = grpcServer
	return server, nil
}

func (server *Server) Start(ctx context.Context) error {
	addr := fmt.Sprintf("127.0.0.1:%d", server.port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	go func() {
		log.Println("starting server at", addr)
		err := server.grpcServer.Serve(lis)
		if err != nil {
			log.Fatalf("unable to start server: %v", err)
		}
	}()

	<-ctx.Done()
	server.grpcServer.Stop()

	return nil
}
