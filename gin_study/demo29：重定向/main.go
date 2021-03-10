package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// HTTP 重定向
	r.GET("/test", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://www.xueyuanjun.com/")
	})

	// 路由重定向
	r.GET("/test1", func(c *gin.Context) {
		c.Request.URL.Path = "/test2"
		r.HandleContext(c)
	})
	r.GET("/test2", func(c *gin.Context) {
		c.JSON(200, gin.H{"hello": "world"})
	})

	r.Run()
}
