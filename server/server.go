package server

import (
	"context"
	"github.com/chemonoworld/grpc-go-demo/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

type greeter struct {
	proto.UnimplementedGreeterServer
}

func (s *greeter) SayHello(ctx context.Context, in *proto.HelloRequest) (*proto.HelloResponse, error) {
	return &proto.HelloResponse{Message: "Hello, " + in.Name + "!"}, nil
}

func (s *greeter) mustEmbedUnimplementedGreeterServer() {}

func Start() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, &greeter{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
