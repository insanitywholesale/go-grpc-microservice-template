package models

import (
	pb "gitlab.com/insanitywholesale/go-grpc-microservice-template/proto/v1"
)

type HelloRepo interface {
	StoreHello(*pb.HelloResponse) error
}
