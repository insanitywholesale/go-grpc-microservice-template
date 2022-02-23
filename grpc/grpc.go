package grpc

import (
	pb "gitlab.com/insanitywholesale/go-grpc-microservice-template/proto/v1"
)

type Server struct {
	pb.UnimplementedHelloServiceServer
}
