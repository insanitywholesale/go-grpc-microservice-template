package main

import (
	"log"
	"net"
	"net/http"

	hellogrpcv1 "gitlab.com/insanitywholesale/go-grpc-microservice-template/grpc/v1"
	hellogrpcv2 "gitlab.com/insanitywholesale/go-grpc-microservice-template/grpc/v2"
	"gitlab.com/insanitywholesale/go-grpc-microservice-template/openapiv2"
	pbv1 "gitlab.com/insanitywholesale/go-grpc-microservice-template/proto/v1"
	pbv2 "gitlab.com/insanitywholesale/go-grpc-microservice-template/proto/v2"
	"gitlab.com/insanitywholesale/go-grpc-microservice-template/rest"
	"gitlab.com/insanitywholesale/go-grpc-microservice-template/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Function to set up and start the gRPC server
func createGRPCServer(listener net.Listener) *grpc.Server {
	// Choose hellorepo v1
	dbv1, err := utils.ChooseRepoV1()
	if err != nil {
		log.Fatal("Failed creating repository v1:", err)
	}
	// Choose hellorepo v2
	dbv2, err := utils.ChooseRepoV2()
	if err != nil {
		log.Fatal("Failed creating repository v2:", err)
	}
	// Create gRPC server
	grpcServer := grpc.NewServer()
	// Register v1 HelloService gRPC server
	pbv1.RegisterHelloServiceServer(grpcServer, hellogrpcv1.Server{DB: dbv1})
	// Register v1 HelloService gRPC server
	pbv2.RegisterHelloServiceServer(grpcServer, hellogrpcv2.Server{DB: dbv2})
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
	docsHandlerv1, err := openapiv2.CreateDocsHandlerV1()
	if err != nil {
		log.Fatal("Failed creating docs handler v1:", err)
	}
	docsHandlerv2, err := openapiv2.CreateDocsHandlerV2()
	if err != nil {
		log.Fatal("Failed creating docs handler v2:", err)
	}

	// Create router
	mux := http.NewServeMux()
	mux.Handle("/api/", restHandler)
	mux.Handle("/docs/v1/", http.StripPrefix("/docs/v1", docsHandlerv1))
	mux.Handle("/docs/v2/", http.StripPrefix("/docs/v2", docsHandlerv2))
	mux.Handle("/docs/", http.StripPrefix("/docs/", docsHandlerv2))

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
	// Create and start the gRPC API server
	grpcServer := createGRPCServer(grpcListener)
	go func() {
		err := grpcServer.Serve(grpcListener)
		if err != nil {
			log.Fatal("Failed starting gRPC server: %w", err)
		}
	}()
	// Create and start the REST API server
	// colon is prepended because we are supposed to pass the entire endpoint
	restServer := createRESTServer(":"+grpcPort, restListener)
	defer func() {
		err := restServer.Serve(restListener)
		if err != nil {
			log.Fatal("Failed starting HTTP server: %w", err)
		}
	}()
}
