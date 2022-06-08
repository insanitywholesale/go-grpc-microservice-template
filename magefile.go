//go:build mage
// +build mage

package main

import (
	"os"

	"github.com/magefile/mage/sh"
)

func CheckFormat() error {
	return sh.Run("diff", "-u", "<(echo -n)", "<(gofmt -d ./)")
}

func Build() error {
	if err := CheckFormat(); err != nil {
		return err
	}
	return sh.Run("go", "install", "-v", "./...")
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

func Protos() error {
	err := sh.Run("buf", "generate", "--timeout=5m30s")
	if err != nil {
		dir, e := os.Getwd()
		if e != nil {
			return e
		}
		return sh.Run("docker", "run", "--rm",
			"-v", dir+":/src",
			"-w", "/src",
			"bufbuild/buf:latest", "generate", "--timeout=5m30s",
		)
	}
	return err
}

func ProtocProtos() error {
	return sh.Run("protoc",
		"-I", "./proto/",
		"-I", "third_party/googleapis",
		"-I", "third_party/grpc-gateway",
		"--go_out=./proto",
		"--go_opt=paths=source_relative",
		"--go-grpc_out=./proto",
		"--go-grpc_opt=paths=source_relative",
		"--openapiv2_out=./openapiv2",
		"--openapiv2_opt=logtostderr=true",
		"--grpc-gateway_out=./proto",
		"--grpc-gateway_opt=logtostderr=true",
		"--grpc-gateway_opt=paths=source_relative",
		"--grpc-gateway_opt=generate_unbound_methods=true",
		"proto/v1/*.proto",
	)
}

func GoRelease() error {
	return sh.Run(
		"go install -v github.com/goreleaser/goreleaser@latest",
		"goreleaser --snapshot --skip-publish --rm-dist",
	)
}
