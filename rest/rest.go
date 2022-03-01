package rest

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	"gitlab.com/insanitywholesale/go-grpc-microservice-template/openapiv2"
	gw "gitlab.com/insanitywholesale/go-grpc-microservice-template/proto/v1"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"io/fs"
	"mime"
	"net/http"
)

// From: https://github.com/johanbrandhorst/grpc-gateway-boilerplate/blob/930554159e8c509132ae7072a5647ac4f7d9e43a/gateway/gateway.go
func CreateDocsHandler() (http.Handler, error) {
	mime.AddExtensionType(".svg", "image/svg+xml")
	// Use subdirectory in embedded files
	subFS, err := fs.Sub(openapiv2.OpenAPIDocsV1, "v1")
	if err != nil {
		return nil, err
	}
	// TODO: Works for /docs/ but doesn't work for /docs so fix it
	return http.StripPrefix("/docs", http.FileServer(http.FS(subFS))), nil
}

func CreateGateway(endpoint string) (http.Handler, error) {
	ctx := context.Background()
	//mux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{OrigName: true, EmitDefaults: true})) // for v1 runtime
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
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterHelloServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
	if err != nil {
		return nil, err
	}
	// Add CORS and return the HTTP handler
	return cors.Default().Handler(mux), nil
}
