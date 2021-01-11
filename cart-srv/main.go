package main

import (
	"log"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	"github.com/yaozhuangyanlingyu/gostudy/cart-srv/api"
	"github.com/yaozhuangyanlingyu/gostudy/cart-srv/lib/mysql"
	"github.com/yaozhuangyanlingyu/gostudy/proto/cart"
	"github.com/yaozhuangyanlingyu/micro-srv/loader"
)

func init() {
	// 加载配置文件
	loader.LoadConfig("")

	// 初始化数据库
	mysql.InitDB()
}

func main() {
	// 将服务注册到consul
	consulReg := consul.NewRegistry(
		registry.Addrs(loader.Config.GetString("consul.address")),
	)

	// 创建micro服务
	service := micro.NewService(
		micro.Name(loader.Config.GetString("server.name")),
		micro.Registry(consulReg),
		micro.Address(loader.Config.GetString("server.address")),
	)

	// 解析命令参数
	service.Init()

	// 注册处理函数
	cart.RegisterCartHandler(service.Server(), new(api.Cart))

	// 运行服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
