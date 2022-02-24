package main

import (
	hellogrpc "gitlab.com/insanitywholesale/go-grpc-microservice-template/grpc"
	pbv1 "gitlab.com/insanitywholesale/go-grpc-microservice-template/proto/v1"
	"gitlab.com/insanitywholesale/go-grpc-microservice-template/rest"
	"gitlab.com/insanitywholesale/go-grpc-microservice-template/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

// TODO: make it return grpc.Server
// Function to set up and start the gRPC server
func startGRPCServer(listener net.Listener) {
	// Choose hellorepo
	db := utils.ChooseRepo()
	// Create gRPC server
	grpcServer := grpc.NewServer()
	// Register v1 HelloService gRPC server
	pbv1.RegisterHelloServiceServer(grpcServer, hellogrpc.Server{DB: db})
	// Enable reflection so the API can be discoverable by something like grpcurl
	reflection.Register(grpcServer)
	// Get port from listener and print it
	port, err := utils.PortFromListener(listener)
	if err != nil {
		log.Println("Failed getting gRPC port from listener:", err)
	} else {
		// Log the server starting as well as the port it is listening on
		log.Println("gRPC server starting on port", port)
	}
	// Start the gRPC server
	// If an error is returned, log the error and exit fatally
	log.Fatal(grpcServer.Serve(listener))
}

// TODO: make it return http.Server (also see rest package TODO)
// Function to set up and start the HTTP server
func startRESTServer(grpcPort string, listener net.Listener) {
	// Get port from listener and print it
	port, err := utils.PortFromListener(listener)
	if err != nil {
		log.Println("Failed getting REST port from listener:", err)
	} else {
		// Log the server starting as well as the port it is listening on
		log.Println("REST server starting on port", port)
	}
	// Start the gRPC-gateway / REST server
	// If an error is returned, log the error and exit fatally
	log.Fatal(rest.RunGateway(grpcPort, listener))
}

func main() {
	// Set up the ports that the servers will listen on
	grpcPort, restPort := utils.SetupPorts()
	// Create listener based on grpc port
	grpcListener, err := utils.ListenerFromPort(grpcPort)
	if err != nil {
		log.Fatalf("gRPC listen failed %v", err)
	}
	// Create listener based on rest port
	restListener, err := utils.ListenerFromPort(restPort)
	if err != nil {
		log.Fatalf("REST listen failed %v", err)
	}
	// Start the gRPC API server
	go startGRPCServer(grpcListener)
	// Start the REST API server
	defer startRESTServer(grpcPort, restListener)
}
