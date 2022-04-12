package utils

import (
	"testing"
)

func TestChooseRepoV1(t *testing.T) {
	repo := ChooseRepoV1()
	t.Log("Repo chosen:", repo)
}

func TestChooseRepoV2(t *testing.T) {
	repo := ChooseRepoV2()
	t.Log("Repo chosen:", repo)
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
