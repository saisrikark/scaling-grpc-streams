package main

import (
	"context"
	"flag"
	"log"
	"scaling-grpc-streams/internal/pkg/server"
)

func main() {
	ctx := context.Background()
	port := flag.Int("port", 8080, "port of which server is hosted")

	server, err := server.NewServer(*port)
	if err != nil {
		log.Println("unable to initialise server", err)
		return
	}

	err = server.Start(ctx)
	if err != nil {
		log.Println("unable to start server", err)
	}
}
