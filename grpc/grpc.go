package grpc

import (
	"context"
	pb "gitlab.com/insanitywholesale/go-grpc-microservice-template/proto/v1"
)

type Server struct {
	pb.UnimplementedHelloServiceServer
}

func (Server) SayHello(context.Context, *pb.Empty) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{HelloWord: "Hello World!"}, nil
}

func (Server) SayCustomHello(_ context.Context, hr *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{HelloWord: "Hello " + hr.CustomWord + "!"}, nil
}
