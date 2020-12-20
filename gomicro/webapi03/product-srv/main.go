package main

import (
	"product-srv/ProService"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
)

type ProductRequest struct {
	Size int `form:"size"`
}

/**
 * 商品API
 */
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
		v1.POST("/prods", func(c *gin.Context) {
			var pr ProductRequest
			err := c.Bind(&pr)
			if err != nil || pr.Size <= 0 {
				pr = ProductRequest{Size: 5}
			}
			c.JSON(200, ProService.GetProductList(pr.Size))
		})
	}

	// 启动go micro服务
	server := web.NewService(
		web.Name("product-service"), // 服务名称
		//web.Address(":8081"),     	// 监听IP:端口
		web.Handler(ginRouter),  // 使用gin路由
		web.Registry(consulReg), // 注册到consul
		// 为注册的服务添加Metadata，指定请求协议为http
		web.Metadata(map[string]string{"protocol": "http"}),
	)
	server.Init()

	// 可以使用go run main.go --server_address :8081
	server.Run()
}
