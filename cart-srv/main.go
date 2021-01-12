package main

import (
	"log"

	"github.com/yaozhuangyanlingyu/gostudy/cart-srv/api"
	"github.com/yaozhuangyanlingyu/gostudy/cart-srv/lib/mysql"
	"github.com/yaozhuangyanlingyu/micro-srv/loader"
	"github.com/yaozhuangyanlingyu/micro-srv/proto/cart"
	"github.com/yaozhuangyanlingyu/micro-srv/server"
)

func init() {
	// 加载配置文件
	loader.LoadConfig("")

	// 初始化数据库
	mysql.InitDB()
}

func main() {
	// 创建service
	service := server.NewAplumService(server.AplumServiceConfig{
		ServiceName: loader.Config.GetString("server.name"),
		ServiceAddr: loader.Config.GetString("server.address"),
		ConsulAddr:  loader.Config.GetString("consul.address"),
	})

	// 解析命令参数
	service.Init()

	// 注册处理函数
	cart.RegisterCartHandler(service.Server(), new(api.Cart))

	// 运行服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
