package grpc

import (
	"context"

	models "gitlab.com/insanitywholesale/go-grpc-microservice-template/models/v2"
	pb "gitlab.com/insanitywholesale/go-grpc-microservice-template/proto/v2"
	"google.golang.org/protobuf/types/known/emptypb"
)

// gRPC server struct, all methods need to be implemented on it
type Server struct {
	// Required so unimplemented methods return error instead of causing compilation failure
	// Source: https://stackoverflow.com/questions/69700899/grpc-error-about-missing-an-unimplemented-server-method
	// Source: https://stackoverflow.com/questions/65079032/grpc-with-mustembedunimplemented-method
	pb.UnimplementedHelloServiceServer
	DB models.HelloRepo
}

func (s Server) SayHello(context.Context, *emptypb.Empty) (*pb.HelloResponse, error) {
	hres := &pb.HelloResponse{HelloWord: "Hello World!"}
	err := s.DB.StoreHello(hres)
	if err != nil {
		return nil, err
	}
	return hres, nil
}

func (s Server) SayCustomHello(_ context.Context, hreq *pb.HelloRequest) (*pb.HelloResponse, error) {
	hres := &pb.HelloResponse{HelloWord: "Hello " + hreq.CustomWord + "!"}
	err := s.DB.StoreHello(hres)
	if err != nil {
		return nil, err
	}
	return hres, nil
}
