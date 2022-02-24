package main

import (
	"context"
	pb "gitlab.com/insanitywholesale/go-grpc-microservice-template/proto/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"net"
	"testing"
)

func TestCreateGRPCServer(t *testing.T) {
	const bufsize = 1024 * 1024
	l := bufconn.Listen(bufsize)
	gs := createGRPCServer(l)
	go gs.Serve(l)

	ctx := context.Background()
	conn, err := grpc.DialContext(
		ctx,
		"bufnet",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) { return l.Dial() }),
		grpc.WithInsecure(),
	)
	if err != nil {
		t.Fatalf("failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	client := pb.NewHelloServiceClient(conn)
	res, err := client.SayHello(ctx, &pb.Empty{})
	if err != nil {
		t.Fatal("Failed to get SayHello response:", err)
	}
	t.Log("SayHello response:", res)
}

func TestStartRESTServer(t *testing.T) {
	t.Skip()
}
