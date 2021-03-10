package main

import (
	"context"
	"fmt"

	"demo01.com/proto"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
)

type Greeter struct {
}

func (this *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	rsp.Greeting = "你好，" + req.Name
	return nil
}

func main() {
	// 将服务注册到consul
	consulReg := consul.NewRegistry(
		registry.Addrs("dev03.aplum-inc.com:8500"),
	)

	// 创建新的服务
	service := micro.NewService(
		micro.Name("go.micro.srv.greeter"),
		micro.Registry(consulReg),
	)

	// 初始化，会解析命令行参数
	service.Init()

	// 注册处理器，调用Greeter服务接口处理请求
	proto.RegisterGreeterHandler(service.Server(), new(Greeter))

	// 启动服务
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
