package rest

import (
	"context"
	"io/fs"
	"mime"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	"gitlab.com/insanitywholesale/go-grpc-microservice-template/openapiv2"
	gw "gitlab.com/insanitywholesale/go-grpc-microservice-template/proto/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
)

// From: https://github.com/johanbrandhorst/grpc-gateway-boilerplate/blob/930554159e8c509132ae7072a5647ac4f7d9e43a/gateway/gateway.go
func CreateDocsHandler() (http.Handler, error) {
	err := mime.AddExtensionType(".svg", "image/svg+xml")
	if err != nil {
		return nil, err
	}
	// Use subdirectory in embedded files
	subFS, err := fs.Sub(openapiv2.OpenAPIDocsV1, "v1")
	if err != nil {
		return nil, err
	}
	return http.StripPrefix("/docs/", http.FileServer(http.FS(subFS))), nil
}

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
	err := gw.RegisterHelloServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
	if err != nil {
		return nil, err
	}
	// Add CORS and return the HTTP handler
	return cors.Default().Handler(mux), nil
}
