package main

import (
	"log"
	"net"
	"net/http"

	hellogrpc "gitlab.com/insanitywholesale/go-grpc-microservice-template/grpc"
	pbv1 "gitlab.com/insanitywholesale/go-grpc-microservice-template/proto/v1"
	"gitlab.com/insanitywholesale/go-grpc-microservice-template/rest"
	"gitlab.com/insanitywholesale/go-grpc-microservice-template/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Function to set up and start the gRPC server
func createGRPCServer(listener net.Listener) *grpc.Server {
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
		// Error is non-fatal because listener does not necesssarily include a port
		log.Println("Failed getting gRPC port from listener:", err)
	} else {
		// Log the server starting as well as the port it is listening on
		log.Println("gRPC server starting on port", port)
	}
	// Create the gRPC server
	return grpcServer
}

// Function to set up and start the HTTP server
func createRESTServer(grpcPort string, listener net.Listener) *http.Server {
	// Get port from listener and print it
	port, err := utils.PortFromListener(listener)
	if err != nil {
		// Error is non-fatal because listener does not necesssarily include a port
		log.Println("Failed getting REST port from listener:", err)
	} else {
		// Log the server starting as well as the port it is listening on
		log.Println("REST server starting on port", port)
	}
	// Create the grpc-gateway / REST server
	restHandler, err := rest.CreateGateway(grpcPort)
	if err != nil {
		log.Fatal("Failed creating grpc-gateway:", err)
	}
	docsHandler, err := rest.CreateDocsHandler()
	if err != nil {
		log.Fatal("Failed creating docs handler:", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/api/", restHandler)
	mux.Handle("/docs/", http.StripPrefix("/docs/", docsHandler))

	return &http.Server{Handler: mux}
}

func main() {
	// Set up the ports that the servers will listen on
	grpcPort, restPort := utils.SetupPorts()
	// Create listener based on grpc port
	grpcListener, err := utils.ListenerFromPort(grpcPort)
	if err != nil {
		log.Fatalf("Creating gRPC listener failed: %v", err)
	}
	// Create listener based on rest port
	restListener, err := utils.ListenerFromPort(restPort)
	if err != nil {
		log.Fatalf("Creating REST listener failed: %v", err)
	}
	// Start the gRPC API server
	grpcServer := createGRPCServer(grpcListener)
	// TODO: fix golangci-lint errcheck by using errors channel like ctrlshiftv
	go grpcServer.Serve(grpcListener)
	// Start the REST API server
	// colon is prepended because we are supposed to pass the entire endpoint
	restServer := createRESTServer(":"+grpcPort, restListener)
	// TODO: fix golangci-lint errcheck by using errors channel like ctrlshiftv
	defer restServer.Serve(restListener)
}
