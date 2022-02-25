package rest

import (
	"gitlab.com/insanitywholesale/go-grpc-microservice-template/grpc"
	pb "gitlab.com/insanitywholesale/go-grpc-microservice-template/proto/v1"
	"gitlab.com/insanitywholesale/go-grpc-microservice-template/utils"
	ggrpc "google.golang.org/grpc"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateGateway(t *testing.T) {
	l, shut := utils.CreateRandomListener()
	defer shut()
	s := ggrpc.NewServer()
	pb.RegisterHelloServiceServer(s, grpc.Server{DB: utils.ChooseRepo()})
	go s.Serve(l)

	// Create ResponseRecorder
	rr := httptest.NewRecorder()

	// Create client request (httptest.NewRequest does server request)
	req, err := http.NewRequest("GET", "/api/v1/hello", nil)
	if err != nil {
		t.Fatal("Failed creating HTTP request:", err)
	}

	// Create the grpc-gateway and get back the http.Handler
	h, err := CreateGateway(l.Addr().(*net.TCPAddr).String())
	if err != nil {
		t.Fatal("Creating gateway failed:", err)
	}

	// Run the server and send the request
	h.ServeHTTP(rr, req)

	// Get HTTP status code
	statusCode := rr.Code
	// Check status code
	if statusCode != http.StatusOK {
		t.Fatalf("Handler returned wrong status code: got %v want %v", statusCode, http.StatusOK)
	}

	// Expected response body
	expected := `{"hello_word":"Hello World!","id":0}`
	// Actual response body
	got := rr.Body.String()
	// Check if expectation matches reality
	if got != expected {
		t.Fatalf("handler returned unexpected body: got %v want %v", got, expected)
	}
}
