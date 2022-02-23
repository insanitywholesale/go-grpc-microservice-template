package main

import (
	hellogrpc "gitlab.com/insanitywholesale/go-grpc-microservice-template/grpc"
	pbv1 "gitlab.com/insanitywholesale/go-grpc-microservice-template/proto/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
)

// Global variable for gRPC server port
var grpcPort string

// Function to determine what port the gRPC server will use
func setupPort() {
	// Get value environment variable
	grpcPort = os.Getenv("HELLO_GRPC_PORT")
	// If empty, default to 15200
	if grpcPort == "" {
		grpcPort = "15200"
	}
}

func startGRPCServer() {
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
	log.Println("gRPC started on port:", grpcPort)
	// Start the gRPC server
	// If an error is returned, log the error and exit fatally
	log.Fatal(grpcServer.Serve(listener))
}

func main() {
	// Set up the port the grpc will listen on
	setupPort()
	// Start the gRPC server
	startGRPCServer()
}
