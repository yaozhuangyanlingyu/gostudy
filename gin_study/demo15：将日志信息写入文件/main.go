package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

//Gin 框架默认将日志输出到控制台，要写入指定的日志文件，可以这么做（src/gin-demo/examples/log_file.go）：

func main() {
	// 日志文件不需要颜色
	gin.DisableConsoleColor()

	// 创建日志文件并设置为 gin.DefaultWriter
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	// 如果你需要同时写入日志文件和控制台，可以这么做：
	// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	router.Run(":8080")
}
