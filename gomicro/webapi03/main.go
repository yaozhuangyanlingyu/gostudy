package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
)

func main() {
	// 将服务注册到consul
	consulReg := consul.NewRegistry(
		registry.Addrs("192.168.95.129:8510"),
	)

	// 使用gin路由
	ginRouter := gin.Default()
	ginRouter.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello world")
	})

	// 启动go micro服务
	server := web.NewService(
		web.Name("product-service"), // 服务名称
		web.Address(":8081"),        // 监听IP:端口
		web.Handler(ginRouter),      // 使用gin路由
		web.Registry(consulReg),     // 注册到consul
	)
	server.Run()
}
