package main

import (
	"context"
	pb "gitlab.com/insanitywholesale/go-grpc-microservice-template/proto/v1"
	"gitlab.com/insanitywholesale/go-grpc-microservice-template/utils"
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
	const bufsize = 1024 * 1024
	gl, shut := utils.CreateRandomListener()
	defer shut()
	gs := createGRPCServer(gl)
	go gs.Serve(gl)

	rl, shut := utils.CreateRandomListener()
	defer shut()
	gp, err := utils.PortFromListener(gl)
	if err != nil {
		t.Fatal(err)
	}
	rs := createRESTServer(":"+gp, rl)
	go rs.Serve(rl)
}
