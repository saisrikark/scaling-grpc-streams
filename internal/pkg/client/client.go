package client

import (
	"scaling-grpc-streams/pkg/application"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	application.GreeterClient
	*grpc.ClientConn
}

func NewClient(addr string) (*Client, error) {

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	greeterClient := application.NewGreeterClient(conn)

	client := &Client{
		GreeterClient: greeterClient,
		ClientConn:    conn,
	}

	return client, nil
}

func (client *Client) Close() error {
	return client.ClientConn.Close()
}
