package main

import (
	"context"
	"log"
	"time"

	pb "model-service/protos"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:3000", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	c := pb.NewModelClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	r, err := c.Register(ctx, &pb.RegisterRequest{
		Email:    "test5@mail.com",
		Password: "password",
	})

	// r, err := c.Activate(ctx, &pb.ActivateRequest{
	// 	Id: 11,
	// })

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Response: %v", r.Response)
}
