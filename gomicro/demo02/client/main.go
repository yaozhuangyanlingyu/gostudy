package main

import (
	"context"
	"demo02/proto"
	"fmt"

	"github.com/micro/go-micro"
)

func main() {
	service := micro.NewService(
		micro.Name("product-srv"),
		micro.Version("product-v0.0.1"),
	)
	service.Init()
	productSrv := proto.NewProductClient("product-srv", service.Client())

	rsp, err := productSrv.GetProductByID(context.TODO(), &proto.GetProductByIDRequest{
		Pid: 8888,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rsp.Code, rsp.Msg, rsp.Data)
}
