package main

import (
	"context"
	"fmt"
	"github.com/noaykkk/grpc-go/pb/hello"
	"google.golang.org/grpc"
	"net"
)

type server struct {
	hello.UnimplementedHelloGRPCServer
}

func (s *server) SayHi(ctx context.Context, req *hello.Req) (res *hello.Res, err error) {
	fmt.Println(req.GetMessage())
	return &hello.Res{Message: "pong"}, nil
}

func hmain() {
	listen, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println(err)
	}
	s := grpc.NewServer()
	hello.RegisterHelloGRPCServer(s, &server{})
	err = s.Serve(listen)
	if err != nil {
		fmt.Println(err)
	}
}
