package openapiv2_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	grpcv1 "gitlab.com/insanitywholesale/go-grpc-microservice-template/grpc/v1"
	grpcv2 "gitlab.com/insanitywholesale/go-grpc-microservice-template/grpc/v2"
	. "gitlab.com/insanitywholesale/go-grpc-microservice-template/openapiv2"
	pbv1 "gitlab.com/insanitywholesale/go-grpc-microservice-template/proto/v1"
	pbv2 "gitlab.com/insanitywholesale/go-grpc-microservice-template/proto/v2"
	"gitlab.com/insanitywholesale/go-grpc-microservice-template/utils"
	ggrpc "google.golang.org/grpc"
)

func TestCreateDocsHandlerV1(t *testing.T) {
	l, shut := utils.CreateRandomListener()
	defer shut()
	s := ggrpc.NewServer()
	repo, _ := utils.ChooseRepoV1()
	pbv1.RegisterHelloServiceServer(s, grpcv1.Server{DB: repo})
	go s.Serve(l)

	// Create ResponseRecorder
	rr := httptest.NewRecorder()

	// Create client request (httptest.NewRequest does server request)
	req, err := http.NewRequest("GET", "/hello.swagger.json", nil)
	if err != nil {
		t.Fatal("Failed creating HTTP request:", err)
	}

	// Create the docs handler and get back the http.Handler
	h, err := CreateDocsHandlerV1()
	if err != nil {
		t.Fatal("Creating docs handler v1 failed:", err)
	}

	// Run the server and send the request
	h.ServeHTTP(rr, req)

	// Get HTTP status code
	statusCode := rr.Code
	// Check if status code is OK
	if statusCode != http.StatusOK {
		t.Fatalf("Handler returned wrong status code: got %v want %v", statusCode, http.StatusOK)
	}
}

func TestCreateDocsHandlerV2(t *testing.T) {
	l, shut := utils.CreateRandomListener()
	defer shut()
	s := ggrpc.NewServer()
	repo, _ := utils.ChooseRepoV2()
	pbv2.RegisterHelloServiceServer(s, grpcv2.Server{DB: repo})
	go s.Serve(l)

	// Create ResponseRecorder
	rr := httptest.NewRecorder()

	// Create client request (httptest.NewRequest does server request)
	req, err := http.NewRequest("GET", "/hello.swagger.json", nil)
	if err != nil {
		t.Fatal("Failed creating HTTP request:", err)
	}

	// Create the docs handler and get back the http.Handler
	h, err := CreateDocsHandlerV2()
	if err != nil {
		t.Fatal("Creating docs handler v2 failed:", err)
	}

	// Run the server and send the request
	h.ServeHTTP(rr, req)

	// Get HTTP status code
	statusCode := rr.Code
	// Check if status code is OK
	if statusCode != http.StatusOK {
		t.Fatalf("Handler returned wrong status code: got %v want %v", statusCode, http.StatusOK)
	}
}
