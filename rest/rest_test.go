package rest

import (
	"encoding/json"
	"gitlab.com/insanitywholesale/go-grpc-microservice-template/grpc"
	models "gitlab.com/insanitywholesale/go-grpc-microservice-template/models/v1"
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

	// Initialize empty MyHello
	mh := &models.MyHello{}
	// Response body
	res := rr.Body.String()
	t.Log("Response body:", res)
	// Check if response body can be marshalled into MyHello
	err = json.NewDecoder(rr.Body).Decode(mh)
	if err != nil {
		t.Fatal("Failed unmarshalling response into MyHello:", err)
	}
	t.Log("Unmarshalled response:", mh)
}
