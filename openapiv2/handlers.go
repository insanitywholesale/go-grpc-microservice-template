package openapiv2

import (
	"io/fs"
	"mime"
	"net/http"
)

// Adapted from: https://github.com/johanbrandhorst/grpc-gateway-boilerplate/blob/930554159e8c509132ae7072a5647ac4f7d9e43a/gateway/gateway.go
func CreateDocsHandlerV1() (http.Handler, error) {
	err := mime.AddExtensionType(".svg", "image/svg+xml")
	if err != nil {
		return nil, err
	}
	// Use subdirectory in embedded files
	subFS, err := fs.Sub(OpenAPIDocsV1, "v1")
	if err != nil {
		return nil, err
	}
	return http.FileServer(http.FS(subFS)), nil
}

func CreateDocsHandlerV2() (http.Handler, error) {
	err := mime.AddExtensionType(".svg", "image/svg+xml")
	if err != nil {
		return nil, err
	}
	// Use subdirectory in embedded files
	subFS, err := fs.Sub(OpenAPIDocsV2, "v2")
	if err != nil {
		return nil, err
	}
	return http.FileServer(http.FS(subFS)), nil
}
