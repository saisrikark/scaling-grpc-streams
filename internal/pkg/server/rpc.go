package server

import (
	"context"
	"fmt"
	"log"
	"scaling-grpc-streams/pkg/application"
	"strings"
	"time"
)

func (server *Server) SayHello(ctx context.Context, helloRequest *application.HelloRequest) (*application.HelloReply, error) {
	return &application.HelloReply{Message: helloRequest.Name}, nil
}

func (server *Server) SayHellos(stream application.Greeter_SayHellosServer) error {
	ctx := context.Background()

	for {
		msg, err := stream.Recv()
		if err != nil {
			return err
		} else {
			log.Println("Received", msg)
		}

		count := strings.Split(msg.Name, "-")[1]

		server.doSharedResourceTask(ctx)

		message := fmt.Sprintf("RESPONSE-%v", count)
		if err := stream.Send(&application.HelloReply{
			Message: message,
		}); err != nil {
			return err
		} else {
			log.Println("Sent", message)
		}
	}
}

func (server *Server) doSharedResourceTask(ctx context.Context) {
	time.Sleep(time.Second * 10)
	// TODOS ensure that only x number of simultaneous access can happen to this function
}
