package utils

import (
	"testing"
)

func TestChooseRepo(t *testing.T) {
	repo := ChooseRepo()
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
}

func TestPortFromListener(t *testing.T) {
	// TODO: skip for now due to problem
	t.Skip()
	// TODO: fix this, causes panic if the check for nil in PortFromListener is not present
	l, _ := ListenerFromPort("1984")
	port, err := PortFromListener(l)
	if err != nil {
		t.Fatal("Problem extracting port from listener:", err)
	}
	t.Log("Port extracted:", port)
}
