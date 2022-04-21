package rest

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	gwv1 "gitlab.com/insanitywholesale/go-grpc-microservice-template/proto/v1"
	gwv2 "gitlab.com/insanitywholesale/go-grpc-microservice-template/proto/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
)

func CreateGateway(endpoint string) (http.Handler, error) {
	ctx := context.Background()
	mux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.HTTPBodyMarshaler{
			Marshaler: &runtime.JSONPb{
				MarshalOptions: protojson.MarshalOptions{
					UseProtoNames:   true,
					EmitUnpopulated: true,
				},
				UnmarshalOptions: protojson.UnmarshalOptions{
					DiscardUnknown: true,
				},
			},
		}),
	)
	// Create a client connection to the gRPC server
	// The gateway acts as a client
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := gwv1.RegisterHelloServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
	if err != nil {
		return nil, err
	}
	err = gwv2.RegisterHelloServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
	if err != nil {
		return nil, err
	}
	// Add CORS and return the HTTP handler
	return cors.Default().Handler(mux), nil
}
