package rest

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/rs/cors"
	gw "gitlab.com/insanitywholesale/go-grpc-microservice-template/proto/v1"
	"google.golang.org/grpc"
	"net/http"
)

func CreateGateway(endpoint string) (http.Handler, error) {
	ctx := context.Background()
	mux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{OrigName: true, EmitDefaults: true}))
	// Create a client connection to the gRPC server
	// The gateway acts as a client
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterHelloServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
	if err != nil {
		return nil, err
	}
	// Add CORS and return the HTTP handler
	return cors.Default().Handler(mux), nil
}
