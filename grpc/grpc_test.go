package grpc

import (
	"context"
	pb "gitlab.com/insanitywholesale/go-grpc-microservice-template/proto/v1"
	"gitlab.com/insanitywholesale/go-grpc-microservice-template/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"net"
	"testing"
)

func TestSayHello(t *testing.T) {
	const bufsize = 1024 * 1024
	l := bufconn.Listen(bufsize)
	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, Server{DB: utils.ChooseRepo()})
	go s.Serve(l)

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
		t.Fatalf("failed to get SayHello response %v", err)
	}
	t.Log("response", res)
}

func TestSayCustomHello(t *testing.T) {
	const bufsize = 1024 * 1024
	l := bufconn.Listen(bufsize)
	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, Server{DB: utils.ChooseRepo()})
	go s.Serve(l)

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
	res, err := client.SayCustomHello(ctx, &pb.HelloRequest{CustomWord: "Test"})
	if err != nil {
		t.Fatalf("failed to get SayCustomHello response %v", err)
	}
	t.Log("response", res)
}