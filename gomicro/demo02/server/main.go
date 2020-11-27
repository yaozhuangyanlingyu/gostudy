package main

import (
	"context"
	"demo02/proto"
	"fmt"
	"log"

	"github.com/micro/go-micro"
)

type Product struct{}

func (_this *Product) GetProductByID(ctx context.Context, req *proto.GetProductByIDRequest, res *proto.GetProductByIDResponse) error {
	fmt.Println(req.Pid)
	res.Code = 200
	res.Msg = "Hello world"
	res.Data = &proto.GetProductByIDResponseData{
		Pid:  req.Pid,
		Name: "玫瑰色连衣裙",
	}
	return nil
}

func main() {
	// Create a new service. Optionally include some options here.
	// 创建新服务.可选择在此处包含一些选项
	service := micro.NewService(
		micro.Name("product-srv"),
		micro.Version("product-v0.0.1"),
	)

	// Init will parse the command line flags.
	// 解析命令行标志
	service.Init()

	// Register handler
	// 注册处理函数
	proto.RegisterProductHandler(service.Server(), new(Product))

	// Run the server
	// 运行服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
