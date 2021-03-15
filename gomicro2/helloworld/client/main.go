package main

import (
	"context"
	"fmt"

	"demo01.com/proto"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
)

func main() {
	// 将服务注册到consul
	consulReg := consul.NewRegistry(
		registry.Addrs("dev03.aplum-inc.com:8500"),
	)

	// 创建一个新服务
	service := micro.NewService(
		micro.Name("go.micro.srv.greeter.client"),
		micro.Registry(consulReg),
	)

	// 初始化
	service.Init()

	// 创建Greeter客户端
	greeter := proto.NewGreeterService("go.micro.srv.greeter", service.Client())

	// 远程调用Greeter服务的Hello方法
	rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{Name: "党静姣"})
	if err != nil {
		fmt.Println(err)
		return
	}

	// print rsp
	fmt.Println(rsp.Greeting)
}
