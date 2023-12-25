package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"scaling-grpc-streams/internal/pkg/client"
	"scaling-grpc-streams/pkg/application"
	"sync"
	"time"
)

var (
	numberOfClients = flag.Int("number", 1, "number of clients to initialise")
	addr            = flag.String("address", "127.0.0.1:8080", "address of the server to contact")
	intervalTime    = *flag.Duration("interval", time.Second*10, "time between requests")
)

func main() {
	flag.Parse()
	wg := &sync.WaitGroup{}

	log.Printf("\nrunning %d clients\n", *numberOfClients)

	for i := 0; i < *numberOfClients; i++ {
		wg.Add(1)
		go runClient(*addr, wg, i+1)
	}

	wg.Wait()
}

func runClient(address string, pwg *sync.WaitGroup, instance int) {
	var wg sync.WaitGroup
	wg.Add(2)

	ctx := context.Background()
	defer pwg.Done()

	message := fmt.Sprintf("REQUEST-%v", instance)

	client, err := client.NewClient(address)
	if err != nil {
		log.Println("Unable to initialise client")
		return
	}

	stream, err := client.SayHellos(ctx)
	if err != nil {
		log.Fatalln("Unable to initialise stream", err)
	}

	go func() {
		defer wg.Done()
		for {
			msg, err := stream.Recv()
			if err != nil {
				log.Println("Error when receiving", err)
			} else {
				log.Println("Got", msg)
			}
		}
	}()

	go func() {
		defer wg.Done()
		for {
			time.Sleep(intervalTime)
			if err := stream.Send(&application.HelloRequest{
				Name: message,
			}); err != nil {
				log.Println("Can't send", err)
			}
		}
	}()

	wg.Wait()
}
