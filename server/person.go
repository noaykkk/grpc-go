package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/noaykkk/grpc-go/pb/person"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type personServe struct {
	person.UnimplementedSearchServiceServer
}

func (*personServe) Search(ctx context.Context, req *person.PersonReq) (*person.PersonRes, error) {
	name := req.GetName()
	res := &person.PersonRes{Name: "verify " + name}
	return res, nil
}

func (*personServe) SearchIn(InServer person.SearchService_SearchInServer) error {
	for {
		req, err := InServer.Recv()
		if err != nil {
			InServer.SendAndClose(&person.PersonRes{Name: "End"})
			break
		}
		fmt.Println(req)
	}
	return nil
}

func (*personServe) SearchOut(req *person.PersonReq, outServer person.SearchService_SearchOutServer) error {
	_cnt := 0
	for {
		if _cnt > 10 {
			break
		}
		time.Sleep(1 * time.Second)
		err := outServer.Send(&person.PersonRes{Name: "I got " + strconv.Itoa(_cnt)})
		if err != nil {
			fmt.Println(err)
		}
		_cnt++
	}
	return nil
}

func (*personServe) SearchIO(ioServer person.SearchService_SearchIOServer) error {
	_cnt := 0
	buffer := make(chan string)
	go func() {
		for {
			if _cnt > 10 {
				buffer <- "End"
				break
			}
			req, err := ioServer.Recv()
			if err != nil {
				buffer <- "End"
				break
			}
			buffer <- req.Name
			_cnt++
		}
	}()
	for {
		s := <-buffer
		if s == "End" {
			break
		}
		err := ioServer.Send(&person.PersonRes{Name: <-buffer})
		if err != nil {
			fmt.Println(err)
			break
		}
	}
	return nil
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go registerGateway(&wg)
	go registerGRPC(&wg)
	wg.Wait()
}

func registerGateway(wg *sync.WaitGroup) {
	conn, _ := grpc.DialContext(
		context.Background(),
		"localhost:8001",
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)

	mux := runtime.NewServeMux()
	gwServer := &http.Server{
		Handler: mux,
		Addr:    ":8090",
	}
	err := person.RegisterSearchServiceHandler(context.Background(), mux, conn)
	if err != nil {
		return
	}
	err = gwServer.ListenAndServe()
	if err != nil {
		return
	}
	wg.Done()
}

func registerGRPC(wg *sync.WaitGroup) {
	listen, err := net.Listen("tcp", ":8001")
	if err != nil {
		return
	}
	context.Background()
	s := grpc.NewServer()
	person.RegisterSearchServiceServer(s, &personServe{})
	s.Serve(listen)
	wg.Done()
}
