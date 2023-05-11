package main

import (
	"context"
	"fmt"
	hello_grpc "github.com/noaykkk/grpc-go/pb"
	"google.golang.org/grpc"
)

func main() {
	dial, err := grpc.Dial("localhost:8000", grpc.WithInsecure())
	defer dial.Close()
	if err != nil {
		fmt.Println(err)
	}
	client := hello_grpc.NewHelloGRPCClient(dial)
	req, err := client.SayHi(context.Background(), &hello_grpc.Req{Message: "ping"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(req.GetMessage())
}
