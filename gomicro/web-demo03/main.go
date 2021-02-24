package main

import (
    "github.com/gin-gonic/gin"
    "github.com/micro/go-micro/registry"
    "github.com/micro/go-micro/web"
    "github.com/micro/go-plugins/registry/consul"
)

// 3.(服务注册)快速把服务注册到Consul中
// 安装go-plugins
// go get -v github.com/micro/go-plugins,原来go-micro consul的支持已经迁移到了go-plugins里面
func main() {
	// 新建一个consul注册的地址，也就是我们consul服务启动的机器ip+端口
    consulReg := consul.NewRegistry(
        registry.Addrs("dev03.aplum-inc.com:8500"),
    )

    ginRouter := gin.Default()
    ginRouter.Handle("GET", "/user", func(context *gin.Context) {
        context.String(200, "user api")
    })

    ginRouter.Handle("GET", "/news", func(context *gin.Context) {
        context.String(200, "news api")
    })

	// go-micro很灵性的实现了注册和反注册，我们启动后直接ctrl+c退出这个server，它会自动帮我们实现反注册
    server := web.NewService( 
        web.Name("test-srv-yaoxf"), 	// 注册进consul服务中的service名字
        //web.Address(":8081"), 			// 注册进consul服务中的端口
        web.Handler(ginRouter), 		// web.Handler()返回一个Option，我们直接把ginRouter穿进去，就可以和gin完美的结合
        web.Registry(consulReg),		// 注册到哪个服务器伤的consul中
    )
    server.Init()
    server.Run()
}