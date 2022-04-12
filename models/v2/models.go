package models

import (
	pb "gitlab.com/insanitywholesale/go-grpc-microservice-template/proto/v2"
)

type HelloRepo interface {
	StoreHello(*pb.HelloResponse) error
}
