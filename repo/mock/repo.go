package mock

import (
	pb "gitlab.com/insanitywholesale/go-grpc-microservice-template/proto/v1"
)

// Variable for hello request id
var helloId uint32 = 0

// Variable for an empty HelloResponse
var emptyHelloResponse = &pb.HelloResponse{}

// Type implementing the HelloRepo interface
type helloRepo []*pb.HelloResponse

// Variable of above type initialized with an empty slice
var hellos helloRepo = []*pb.HelloResponse{}

// Function returning the initialized variable implementing HelloRepo interface
func NewMockRepo() (helloRepo, error) {
	return hellos, nil
}

// Implementation of StoreHello from HelloRepo interface
func (helloRepo) StoreHello(hr *pb.HelloResponse) error {
	hr.Id = helloId
	helloId++
	hellos = append(hellos, hr)
	return nil
}
