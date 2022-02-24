package main

import (
	"testing"
)

func TestChooseRepo(t *testing.T) {
	repo := chooseRepo()
	t.Log("Repo chosen:", repo)
}

func TestSetupPorts(t *testing.T) {
	grpcport, restport := setupPorts()
	if grpcport != "15200" {
		t.Fatal("Problem setting up gRPC port")
	}
	if restport != "8080" {
		t.Fatal("Problem setting up REST port")
	}
}

func TestStartGRPCServer(t *testing.T) {
}

func TestStartRESTServer(t *testing.T) {
}
