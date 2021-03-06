package grpc_test

import (
	"context"
	"net"
	"testing"

	. "gitlab.com/insanitywholesale/go-grpc-microservice-template/grpc/v2"
	pb "gitlab.com/insanitywholesale/go-grpc-microservice-template/proto/v2"
	"gitlab.com/insanitywholesale/go-grpc-microservice-template/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestSayHello(t *testing.T) {
	const bufsize = 1024 * 1024
	l := bufconn.Listen(bufsize)
	s := grpc.NewServer()
	repo, _ := utils.ChooseRepoV2()
	pb.RegisterHelloServiceServer(s, Server{DB: repo})
	go s.Serve(l)

	ctx := context.Background()
	conn, err := grpc.DialContext(
		ctx,
		"bufnet",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) { return l.Dial() }),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		t.Fatalf("failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	client := pb.NewHelloServiceClient(conn)
	res, err := client.SayHello(ctx, &emptypb.Empty{})
	if err != nil {
		t.Fatalf("failed to get SayHello response %v", err)
	}
	t.Log("response", res)
}

func TestSayCustomHello(t *testing.T) {
	const bufsize = 1024 * 1024
	l := bufconn.Listen(bufsize)
	s := grpc.NewServer()
	repo, _ := utils.ChooseRepoV2()
	pb.RegisterHelloServiceServer(s, Server{DB: repo})
	go s.Serve(l)

	ctx := context.Background()
	conn, err := grpc.DialContext(
		ctx,
		"bufnet",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) { return l.Dial() }),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
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
