package main

import (
	"context"
	"fmt"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
    "github.com/yaozhuangyanlingyu/gostudy/cart-srv/proto/cart"
)

func main() {
	// 将服务注册到cansul
	consulReg := consul.NewRegistry(
		registry.Addrs("dev03.aplum-inc.com:8500"),
	)

	// 创建client
	myService := micro.NewService(
		micro.Name("cart-srv-client"),
		micro.Registry(consulReg),
	)
	cartSrv := cart.NewCartService(
		"cart-srv-test",
		myService.Client(),
	)

	// 获取购物车列表数据
	req := cart.ListRequest{
		UserID:         46399747,                 // 用户ID
		KeyID:          "",                       // KeyID
		SourcePlatform: cart.SourcePlatform_PLUM, // 来源平台
	}
	rsp, err := cartSrv.List(context.Background(), &req)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(rsp)
}
