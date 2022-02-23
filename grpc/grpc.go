package grpc

import (
	"context"
	pb "gitlab.com/insanitywholesale/go-grpc-microservice-template/proto/v1"
)

type Server struct {
	// Required so unimplemented methods return error instead of causing compilation failure
	// Source: https://stackoverflow.com/questions/69700899/grpc-error-about-missing-an-unimplemented-server-method
	// Source: https://stackoverflow.com/questions/65079032/grpc-with-mustembedunimplemented-method
	pb.UnimplementedHelloServiceServer
}

func (Server) SayHello(context.Context, *pb.Empty) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{HelloWord: "Hello World!"}, nil
}

func (Server) SayCustomHello(_ context.Context, hr *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{HelloWord: "Hello " + hr.CustomWord + "!"}, nil
}
