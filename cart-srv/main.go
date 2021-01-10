package main

import (
	"context"
	"log"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	"github.com/yaozhuangyanlingyu/gostudy/proto/cart"
)

type Cart struct{}

func (_this *Cart) List(ctx context.Context, req *cart.ListRequest, rsp *cart.ListResponse) error {
	rsp.Code = 200
	rsp.Msg = "success"
	return nil
}

func main() {
	// 将服务注册到consul
	consulReg := consul.NewRegistry(
		registry.Addrs("dev03.aplum-inc.com:8500"),
	)

	// 创建micro服务
	service := micro.NewService(
		micro.Name("cart-srv-test"),
		micro.Registry(consulReg),
	)

	// 解析命令参数
	service.Init()

	// 注册处理函数
	cart.RegisterCartHandler(service.Server(), new(Cart))

	// 运行服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
