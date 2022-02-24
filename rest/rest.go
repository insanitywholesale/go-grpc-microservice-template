package rest

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/rs/cors"
	gw "gitlab.com/insanitywholesale/go-grpc-microservice-template/proto/v1"
	"google.golang.org/grpc"
	"net"
	"net/http"
)

// TODO: make it return http.Server{Handler: handler} and listener
func RunGateway(grpcport string, listener net.Listener) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{OrigName: true, EmitDefaults: true}))
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterHelloServiceHandlerFromEndpoint(ctx, mux, ":"+grpcport, opts)
	if err != nil {
		return err
	}
	handler := cors.Default().Handler(mux)
	return http.Serve(listener, handler)
}
