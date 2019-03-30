package main

import (
	"context"
	"log"
	"net"

	pb "model-service/proto"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Create(ctx context.Context, in *pb.CreateRequest) (*pb.CreateReply, error) {
	return &pb.CreateReply{Message: "Hello again " + in.Name}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":3000")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterModelServer(s, &server{})

	err = s.Serve(listener)

	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
}
