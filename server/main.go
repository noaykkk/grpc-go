package main

import (
	"context"
	"fmt"
	hello_grpc "github.com/noaykkk/grpc-go/pb"
	"google.golang.org/grpc"
	"net"
)

type server struct {
	hello_grpc.UnimplementedHelloGRPCServer
}

func (s *server) SayHi(ctx context.Context, req *hello_grpc.Req) (res *hello_grpc.Res, err error) {
	fmt.Println(req.GetMessage())
	return &hello_grpc.Res{Message: "pong"}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println(err)
	}
	s := grpc.NewServer()
	hello_grpc.RegisterHelloGRPCServer(s, &server{})
	err = s.Serve(listen)
	if err != nil {
		fmt.Println(err)
	}
}
