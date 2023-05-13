package main

import (
	"context"
	"fmt"
	"github.com/noaykkk/grpc-go/pb/person"
	"google.golang.org/grpc"
	"strconv"
	"time"
)

func main() {
	dial, err := grpc.Dial("localhost:8001", grpc.WithInsecure())
	if err != nil {
		return
	}
	client := person.NewSearchServiceClient(dial)

	// For Search()
	//res, err := client.Search(context.Background(), &person.PersonReq{Name: "admin"})
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(res)

	// For SearchIn()
	in, err := client.SearchIn(context.Background())
	if err != nil {
		return
	}
	_cnt := 0
	for {
		if _cnt > 10 {
			res, _ := in.CloseAndRecv()
			fmt.Println(res)
			break
		}
		time.Sleep(1 * time.Second)
		err := in.Send(&person.PersonReq{Name: "Input " + strconv.Itoa(_cnt)})
		if err != nil {
			fmt.Println(err)
			break
		}
		_cnt++
	}

	// For SearchOut()
	//c, _ := client.SearchOut(context.Background(), &person.PersonReq{Name: "admin"})
	//for {
	//	req, err := c.Recv()
	//	if err != nil {
	//		fmt.Println(err)
	//		break
	//	}
	//	fmt.Println(req)
	//}

	// For SearchIO()
	//c, _ := client.SearchIO(context.Background())
	//wg := sync.WaitGroup{}
	//wg.Add(2)
	//go func() {
	//	for {
	//		err := c.Send(&person.PersonReq{Name: "admin"})
	//		if err != nil {
	//			wg.Done()
	//			break
	//		}
	//		time.Sleep(1 * time.Second)
	//	}
	//}()
	//go func() {
	//	for {
	//		req, err := c.Recv()
	//		if err != nil {
	//			fmt.Println(err)
	//			wg.Done()
	//			break
	//		}
	//		fmt.Println(req)
	//	}
	//}()
	//wg.Wait()
}
