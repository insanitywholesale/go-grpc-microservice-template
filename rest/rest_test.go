package rest_test

import (
	"encoding/json"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"

	grpcv1 "gitlab.com/insanitywholesale/go-grpc-microservice-template/grpc/v1"
	grpcv2 "gitlab.com/insanitywholesale/go-grpc-microservice-template/grpc/v2"
	pbv1 "gitlab.com/insanitywholesale/go-grpc-microservice-template/proto/v1"
	pbv2 "gitlab.com/insanitywholesale/go-grpc-microservice-template/proto/v2"
	. "gitlab.com/insanitywholesale/go-grpc-microservice-template/rest"
	"gitlab.com/insanitywholesale/go-grpc-microservice-template/utils"
	ggrpc "google.golang.org/grpc"
)

type testHelloResponse struct {
	HelloWord string `json:"hello_word"`
	ID        uint32 `json:"id"`
}

func TestCreateGateway(t *testing.T) {
	l, shut := utils.CreateRandomListener()
	defer shut()
	s := ggrpc.NewServer()
	pbv1.RegisterHelloServiceServer(s, grpcv1.Server{DB: utils.ChooseRepoV1()})
	pbv2.RegisterHelloServiceServer(s, grpcv2.Server{DB: utils.ChooseRepoV2()})
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
	t.Log("Got status code 200 from /api/v1/hello")

	// Initialize empty testHelloResponse
	mh := &testHelloResponse{}
	// Response body
	res := rr.Body.String()
	t.Log("Response body:", res)
	// Check if response body can be marshalled into testHelloResponse
	err = json.NewDecoder(rr.Body).Decode(mh)
	if err != nil {
		t.Fatal("Failed unmarshalling response into testHelloResponse:", err)
	}
	t.Log("Unmarshalled response:", mh)
}
