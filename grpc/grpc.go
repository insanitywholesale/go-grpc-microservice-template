package grpc

import (
	pb "gitlab.com/insanitywholesale/lister/proto/v1"
)

type Server struct {
	pb.UnimplementedHelloServiceServer
}
