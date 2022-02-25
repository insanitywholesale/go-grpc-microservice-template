.PHONY: checkformat getdeps protos gorelease

checkformat:
	diff -u <(echo -n) <(gofmt -d ./)

getdeps:
	/bin/sh -c 'type protoc'
	export GO111MODULE=on
	go install -v google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install -v google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install -v github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	go install -v github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest

generate:
	docker run -v $$(pwd):/src -w /src --rm bufbuild/buf:latest generate

protos:
	protoc -I ./proto/ -I third_party/googleapis -I third_party/grpc-gateway \
	--go_out=./proto \
	--go_opt=paths=source_relative \
	--go-grpc_out=./proto \
	--go-grpc_opt=paths=source_relative \
	--openapiv2_out=./openapiv2 \
	--openapiv2_opt=logtostderr=true \
	--grpc-gateway_out=./proto \
	--grpc-gateway_opt=logtostderr=true \
	--grpc-gateway_opt=paths=source_relative \
	--grpc-gateway_opt=generate_unbound_methods=true \
	proto/v1/*.proto

gorelease:
	go install -v github.com/goreleaser/goreleaser@latest
	goreleaser --snapshot --skip-publish --rm-dist
