package main

import (
	"context"
	"fmt"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
    "github.com/yaozhuangyanlingyu/micro-srv/proto/cart"
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

    // 列表
    //_list(cartSrv)

    // 删除
    _delete(cartSrv)
}

// 删除
func _delete(cartSrv cart.CartService) {
    pids := []int32{110483, 1159871}
    req := cart.DeleteRequest{
        UserID : 46399747,
        ProductID : pids,
        KeyID : "",
		SourcePlatform: cart.SourcePlatform_PLUM, // 来源平台
    }
	rsp, err := cartSrv.Delete(context.Background(), &req)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(rsp)
}

// 列表
func _list(cartSrv cart.CartService) {
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
