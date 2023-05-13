package main

import (
	"context"
	"fmt"
	"github.com/noaykkk/grpc-go/pb/hello"
	"google.golang.org/grpc"
)

func hmain() {
	dial, err := grpc.Dial("localhost:8000", grpc.WithInsecure())
	defer dial.Close()
	if err != nil {
		fmt.Println(err)
	}
	client := hello.NewHelloGRPCClient(dial)
	req, err := client.SayHi(context.Background(), &hello.Req{Message: "ping"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(req.GetMessage())
}
