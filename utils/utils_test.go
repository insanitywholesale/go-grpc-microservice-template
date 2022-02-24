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
