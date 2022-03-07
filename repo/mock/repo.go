package mock

import (
	models "gitlab.com/insanitywholesale/go-grpc-microservice-template/models/v1"
	pb "gitlab.com/insanitywholesale/go-grpc-microservice-template/proto/v1"
)

// Variable for hello request id
var helloID uint32 = 1

// Type implementing the HelloRepo interface
type helloRepo []*pb.HelloResponse

// Variable of above type initialized with an empty slice
var hellos helloRepo = []*pb.HelloResponse{}

// Function returning the initialized variable implementing HelloRepo interface
func NewMockRepo() (models.HelloRepo, error) {
	return hellos, nil
}

// Implementation of StoreHello from HelloRepo interface
func (helloRepo) StoreHello(hr *pb.HelloResponse) error {
	hr.Id = helloID
	helloID++
	hellos = append(hellos, hr)
	return nil
}
