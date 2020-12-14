package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/web"
)

func main() {
	// 使用gin路由
	ginRouter := gin.Default()
	v1 := ginRouter.Group("/v1")
	{
		v1.GET("/index", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"name": "hello world",
			})
		})
	}

	// 启动go micro服务
	server := web.NewService(
		web.Address(":8082"),   // 监听IP:端口
		web.Handler(ginRouter), // 使用gin路由
	)
	server.Run()
}
