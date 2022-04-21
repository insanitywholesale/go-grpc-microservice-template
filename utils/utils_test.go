package utils_test

import (
	"testing"

	. "gitlab.com/insanitywholesale/go-grpc-microservice-template/utils"
)

func TestChooseRepoV1(t *testing.T) {
	repo, err := ChooseRepoV1()
	if err != nil {
		t.Error("Error creating v2 repo:", err)
	}
	t.Log("Repo v1 chosen:", repo)
}

func TestChooseRepoV2(t *testing.T) {
	repo, err := ChooseRepoV2()
	if err != nil {
		t.Error("Error creating v2 repo:", err)
	}
	t.Log("Repo v2 chosen:", repo)
}

func TestSetupPorts(t *testing.T) {
	grpcport, restport := SetupPorts()
	if grpcport != "15200" {
		t.Fatal("Problem setting up gRPC port")
	}
	if restport != "8080" {
		t.Fatal("Problem setting up REST port")
	}
}

func TestListenerFromPort(t *testing.T) {
	l, err := ListenerFromPort("1984")
	if err != nil {
		t.Fatal("Problem creating listener from port:", err)
	}
	t.Log("Listener created:", l)
	l.Close()
}

func TestPortFromListener(t *testing.T) {
	l, err := ListenerFromPort("1984")
	if err != nil {
		t.Fatal("Problem creating listener from port:", err)
	}
	defer l.Close()
	port, err := PortFromListener(l)
	if err != nil {
		t.Fatal("Problem extracting port from listener:", err)
	}
	t.Log("Port extracted:", port)
}

func TestCreateRandomListener(t *testing.T) {
	l, shut := CreateRandomListener()

	port, err := PortFromListener(l)
	if err != nil {
		t.Fatal("Problem extracting port from listener:", err)
	}
	t.Log("Port extracted:", port)

	shut()
	t.Log("Shut() succeeded")
}
