package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/get", func(c *gin.Context) {
		// 从查询字符串获取数据
		ids := c.QueryMap("ids")
		// 从请求表单获取数据
		names := c.PostFormMap("names")

		fmt.Printf("ids: %v; names: %v", ids, names)
	})

	router.Run(":8080")
}

// curl http://60.205.148.21:8080/get?ids[a]=1234&ids[b]=hello
