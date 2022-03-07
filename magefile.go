//go:build mage
// +build mage

package main

import (
	"github.com/magefile/mage/sh"
)

func Build() error {
	if err := sh.Run("diff", "-u", "<(echo -n)", "<(gofmt -d ./)"); err != nil {
		return err
	}
	return sh.Run("go", "install", "./...")
}

func GetDeps() error {
	if err := sh.Run("/bin/sh", "-c", "'type protoc'"); err != nil {
		return err
	}
	return sh.Run("go", "install",
		"google.golang.org/protobuf/cmd/protoc-gen-go@latest",
		"google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest",
		"github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest",
		"github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest",
	)
}
