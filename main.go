package main

import (
	hellogrpc "gitlab.com/insanitywholesale/go-grpc-microservice-template/grpc"
	pbv1 "gitlab.com/insanitywholesale/go-grpc-microservice-template/proto/v1"
	"gitlab.com/insanitywholesale/go-grpc-microservice-template/rest"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
)

// Function to determine what port the gRPC and HTTP servers will use
func setupPorts() (string, string) {
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

// Function to set up and start the gRPC server
func startGRPCServer(grpcPort string) {
	// Create listener on gRPC port
	listener, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		log.Fatalf("listen failed %v", err)
	}

	// Create gRPC server
	grpcServer := grpc.NewServer()
	// Register v1 HelloService gRPC server
	pbv1.RegisterHelloServiceServer(grpcServer, hellogrpc.Server{})
	// Enable reflection so the API can be discoverable by something like grpcurl
	reflection.Register(grpcServer)
	// Log the server starting as well as the port it is listening on
	log.Println("gRPC server starting on port:", grpcPort)
	// Start the gRPC server
	// If an error is returned, log the error and exit fatally
	log.Fatal(grpcServer.Serve(listener))
}

// Function to set up and start the HTTP server
func startRESTServer(grpcPort string, restPort string) {
	log.Println("REST server starting on port", restPort)
	log.Fatal(rest.RunGateway(grpcPort, restPort))
}

func main() {
	// Set up the ports that the servers will listen on
	grpcPort, restPort := setupPorts()
	// Start the gRPC API server
	go startGRPCServer(grpcPort)
	// Start the REST API server
	defer startRESTServer(grpcPort, restPort)
}
