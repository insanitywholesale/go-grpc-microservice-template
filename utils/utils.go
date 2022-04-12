package utils

import (
	"errors"
	"net"
	"os"
	"strconv"
	"strings"

	modelsv1 "gitlab.com/insanitywholesale/go-grpc-microservice-template/models/v1"
	modelsv2 "gitlab.com/insanitywholesale/go-grpc-microservice-template/models/v2"
	mockv1 "gitlab.com/insanitywholesale/go-grpc-microservice-template/repo/v1/mock"
	mockv2 "gitlab.com/insanitywholesale/go-grpc-microservice-template/repo/v2/mock"
)

// Function to select data repository backend v1
func ChooseRepoV1() modelsv1.HelloRepo {
	mockrepo, _ := mockv1.NewMockRepo()
	return mockrepo
}

// Function to select data repository backend v2
func ChooseRepoV2() modelsv2.HelloRepo {
	mockrepo, _ := mockv2.NewMockRepo()
	return mockrepo
}

// Function to determine what port the gRPC and HTTP servers will use
func SetupPorts() (string, string) {
	// Variables for ports
	var grpcPort string
	var restPort string
	// Get value environment variable
	grpcPort = os.Getenv("HELLO_GRPC_PORT")
	// If empty, default to 15200
	if grpcPort == "" {
		grpcPort = "15200"
	}
	restPort = os.Getenv("HELLO_REST_PORT")
	if restPort == "" {
		restPort = "8080"
	}
	return grpcPort, restPort
}

func ListenerFromPort(port string) (net.Listener, error) {
	// Create listener on provided port
	l, err := net.Listen("tcp4", ":"+port)
	if err != nil {
		//TODO: wrap the error to be more explanatory
		return nil, err
	}
	return l, nil
}

func PortFromListener(l net.Listener) (string, error) {
	if l == nil {
		return "", errors.New("provided listener is nil")
	}
	addrSlice := strings.Split(l.Addr().String(), ":")
	port := addrSlice[len(addrSlice)-1]
	_, err := strconv.Atoi(port)
	if err != nil {
		//TODO: wrap the error to be more explanatory
		return "", err
	}
	return port, nil
}

// Function to create listener with random open port for testing purposes
func CreateRandomListener() (l net.Listener, shut func()) {
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		panic(err)
	}

	return l, func() {
		_ = l.Close()
	}
}
