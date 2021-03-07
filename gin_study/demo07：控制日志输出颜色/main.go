package main

import "github.com/gin-gonic/gin"

func main() {
	// 禁止日志颜色
	//gin.DisableConsoleColor()

	// 强制设置日志颜色
	gin.ForceConsoleColor()

	// 使用默认中间件（logger 和 recovery）创建 gin 路由器
	router := gin.Default()

	// 定义路由
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// 启动服务器
	router.Run(":8080")
}
