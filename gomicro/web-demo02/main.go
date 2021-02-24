package main

import (
    "github.com/gin-gonic/gin"
    "github.com/micro/go-micro/web"
)

// 2.引进外部框架gin生成web API
// //web.Handler()返回一个Option，我们直接把ginRouter穿进去，就可以和gin完美的结合
func main() {
	ginRouter := gin.Default()

    ginRouter.Handle("GET", "/user", func(context *gin.Context) {
        context.String(200, "user api")
    })

    ginRouter.Handle("GET", "/news", func(context *gin.Context) {
        context.String(200, "news api")
    })

    server := web.NewService(
        web.Address(":8081"),
        web.Handler(ginRouter),
    )
    server.Run()
}