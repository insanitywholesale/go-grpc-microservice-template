package grpc

import (
	"context"
	models "gitlab.com/insanitywholesale/go-grpc-microservice-template/models/v1"
	pb "gitlab.com/insanitywholesale/go-grpc-microservice-template/proto/v1"
	"gitlab.com/insanitywholesale/go-grpc-microservice-template/repo/mock"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"net"
	"testing"
)

// TODO: should be moved into repo package
func chooseRepo() models.HelloRepo {
	mockrepo, _ := mock.NewMockRepo()
	return mockrepo
}

func TestSayHello(t *testing.T) {
	const bufsize = 1024 * 1024
	l := bufconn.Listen(bufsize)
	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, Server{DB: chooseRepo()})
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
	pb.RegisterHelloServiceServer(s, Server{DB: chooseRepo()})
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
