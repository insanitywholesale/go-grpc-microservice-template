package mock_test

import (
	"testing"

	pb "gitlab.com/insanitywholesale/go-grpc-microservice-template/proto/v1"
	. "gitlab.com/insanitywholesale/go-grpc-microservice-template/repo/v1/mock"
)

func TestNewMockRepo(t *testing.T) {
	repo, err := NewMockRepo()
	if err != nil {
		t.Fatal("Creating new mock repo failed:", err)
	}
	t.Log("New mock repo created", repo)
}

func TestStoreHello(t *testing.T) {
	repo, _ := NewMockRepo()
	err := repo.StoreHello(&pb.HelloResponse{HelloWord: "Hello Test World!"})
	if err != nil {
		t.Fatal("Storing Hello failed:", err)
	}
}
