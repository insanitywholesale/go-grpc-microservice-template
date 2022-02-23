package mock

import (
	pb "gitlab.com/insanitywholesale/go-grpc-microservice-template/proto/v1"
)

// Variable for hello request id
var helloId = 0

// Variable for an empty HelloRequest
var emptyHelloRequest = &pb.HelloRequest{}

// Type implementing the HelloRepo interface
type helloRepo []*pb.HelloRequest

// Variable of above type initialized with an empty slice
var hellos helloRepo = []*pb.HelloRequest{}

// Function returning the initialized variable implementing HelloRepo interface
func NewMockRepo() (helloRepo, error) {
	return hellos, nil
}

// Implementation of StoreHello from HelloRepo interface
func (helloRepo) StoreHello(hr *pb.HelloRequest) error {
	hr.Id = helloId
	helloId++
	hellos = append(hellos, hr)
	return nil
}
