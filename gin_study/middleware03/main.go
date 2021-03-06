package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	// 全局中间件
	// 使用 Logger 中间件
	r.Use(gin.Logger())

	// 使用 Recovery 中间件
	r.Use(gin.Recovery())

	shopGroup := r.Group("/shop")
	shopGroup.Use(AuthRequired(), SetProductCache())
	{
		shopGroup.GET("/index1", indexHandle1)
		shopGroup.GET("/index2", indexHandle2)
	}

	r.Run() // listen and serve on 0.0.0.0:8080
}

func indexHandle1(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "200",
		"msg":    "indexHandle1",
	})
}

func indexHandle2(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "200",
		"msg":    "indexHandle2",
	})
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("AuthRequired 请求URI：", c.Request.RequestURI)
	}
}

func SetProductCache() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("SetProductCache 请求URI：", c.Request.RequestURI)
	}
}
