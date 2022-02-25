package models

import (
	pb "gitlab.com/insanitywholesale/go-grpc-microservice-template/proto/v1"
)

type HelloRepo interface {
	StoreHello(*pb.HelloResponse) error
}

type MyHello struct {
	HelloWord string `json:"hello_word"`
	Id uint32 `json:"id"`
}
