package main

import (
	"context"
	"demo03/proto"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
)

func main() {
	// 将服务注册到cansul
	consulReg := consul.NewRegistry(
		registry.Addrs("dev03.aplum-inc.com:8500"),
	)

	// 使用gin路由
	myService := micro.NewService(micro.Name("product-api.client"))
	productSrv := proto.NewProductService("product-srv-yaoxf", myService.Client())
	ginRouter := gin.Default()
	v1 := ginRouter.Group("/v1")
	{
		v1.GET("/prods", func(c *gin.Context) {
			// 验证参数
			var rquestParams proto.GetProductListRequest
			size, err := strconv.ParseInt(c.DefaultQuery("size", "10"), 10, 32)
			if err != nil {
				c.JSON(500, gin.H{
					"error": err.Error(),
				})
				return
			}
			rquestParams.Size = int32(size)

			// 请求product服务
			rsp, err := productSrv.GetProductList(context.Background(), &rquestParams)
			if err != nil {
				c.JSON(500, gin.H{
					"error": err.Error(),
				})
				return
			}

			// 获取数据
			c.JSON(200, gin.H{
				"size": rquestParams.Size,
				"data": rsp.Data,
			})
		})
	}

	// 启动go micro服务
	server := web.NewService(
		web.Name("product-api"), // 服务名称
		web.Handler(ginRouter),  // 注册gin路由
		web.Registry(consulReg), // 注册到consul

		// 为注册的服务添加Metadata，指定请求协议为http
		web.Metadata(map[string]string{"protocol": "http"}),
	)

	// 解析命令参数
	server.Init()

	// 可以使用go run main.go --server_address :8081
	server.Run()
}
