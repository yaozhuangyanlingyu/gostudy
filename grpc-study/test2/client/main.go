package main

import (
	"context"
	"fmt"

	"github.com/yaozhuangyanlingyu/gostudy/grpc-study/test2/proto"
	"google.golang.org/grpc"
)

func main() {
	// 连接服务器
	conn, err := grpc.Dial(":8848", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("faild to connect: %v", err)
	}
	defer conn.Close()

	c := proto.NewProductClient(conn)
	res, err := c.GetProduct(context.Background(), &proto.ProductRequest{Pid: "123"})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}
