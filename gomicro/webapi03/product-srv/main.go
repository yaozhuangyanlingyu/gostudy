package main

import (
	"product-srv/ProService"

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
	v1 := ginRouter.Group("/v1")
	{
		v1.GET("/prods", func(c *gin.Context) {
			c.JSON(200, ProService.GetProductList(5))
		})
	}

	// 启动go micro服务
	server := web.NewService(
		web.Name("product-service"), // 服务名称
		web.Address(":8081"),        // 监听IP:端口
		web.Handler(ginRouter),      // 使用gin路由
		web.Registry(consulReg),     // 注册到consul
	)
	server.Run()
}
