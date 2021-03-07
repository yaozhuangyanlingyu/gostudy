package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// 定义路由
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// 启动服务器
	//router.Run(":8080")

	// 可以直接使用 http.ListenAndServe() 来配置监听地址和端口以及路由器：
	// http.ListenAndServe(":8080", router)

	// 或者通过下面这种方式来自定义 HTTP 服务器：
	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
