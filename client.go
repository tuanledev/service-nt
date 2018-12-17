package main

import (
	"context"
	"log"
	pb "service-nt/proto"
	"time"

	"google.golang.org/grpc"
)

const (
	address     = "103.1.239.145:50051"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewDeviceServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetData(ctx, &pb.DeviceStatus{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %v , %s", r.Status, r.Time)
}
