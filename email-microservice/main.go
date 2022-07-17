package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "0.0.0.0:8080"
)

func main() {
	// create connection to grpc email service
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grcp.WithBlock())
	if err != nil {
		log.Fatal("Failed to connect: %v", err)
	}
	defer conn.Close()

	// client := pb.
}
