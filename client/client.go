package client

import (
	"context"
	pb "github.com/dangrmous/transit-pdx/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type Client struct {
	clientLogger *log.Logger
}

func New(logger *log.Logger) *Client {
	return &Client{
		clientLogger: logger,
	}
}

func (c *Client) Send() {
	c.clientLogger.Println("Sending message")
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient("localhost:8000", opts...)
	if err != nil {
		c.clientLogger.Println(err)
	}
	client := pb.NewTransitPDXClient(conn)
	result, err := client.GetScheduledTimes(context.Background(), &pb.StopId{Stop_Id: 123})
	if err != nil {
		c.clientLogger.Println(err)
	}
	c.clientLogger.Println(result)
	defer conn.Close()
}
