package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/web"
)

func main() {
	ginRouter := gin.Default()
	ginRouter.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello world")
	})
	server := web.NewService(
		web.Address(":8081"),
		web.Handler(ginRouter),
	)
	server.Run()
}
