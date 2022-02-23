package grpc

import (
	"context"
	models "gitlab.com/insanitywholesale/go-grpc-microservice-template/models/v1"
	pb "gitlab.com/insanitywholesale/go-grpc-microservice-template/proto/v1"
	"gitlab.com/insanitywholesale/go-grpc-microservice-template/repo/mock"
)

var db models.HelloRepo

// TODO: this can likely be implemented in a better way e.g. in Server struct
func init() {
	mockrepo, _ := mock.NewMockRepo()
	db = mockrepo
	return
}

// gRPC server struct, all methods need to be implemented on it
type Server struct {
	// Required so unimplemented methods return error instead of causing compilation failure
	// Source: https://stackoverflow.com/questions/69700899/grpc-error-about-missing-an-unimplemented-server-method
	// Source: https://stackoverflow.com/questions/65079032/grpc-with-mustembedunimplemented-method
	pb.UnimplementedHelloServiceServer
}

func (Server) SayHello(context.Context, *pb.Empty) (*pb.HelloResponse, error) {
	hres := &pb.HelloResponse{HelloWord: "Hello World!"}
	db.StoreHello(hres)
	return hres, nil
}

func (Server) SayCustomHello(_ context.Context, hreq *pb.HelloRequest) (*pb.HelloResponse, error) {
	hres := &pb.HelloResponse{HelloWord: "Hello " + hreq.CustomWord + "!"}
	db.StoreHello(hres)
	return hres, nil
}
