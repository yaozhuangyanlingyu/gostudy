package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/post", func(c *gin.Context) {
		id := c.Query("id")                 // 查询字符串
		page := c.DefaultQuery("page", "0") // 查询字符串（带默认值）
		name := c.PostForm("name")          //  POST 表单数据
		message := c.PostForm("message")    //  同上

		fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
	})

	router.Run(":8080")
}

// curl -X POST -H "Content-Type:application/x-www-form-urlencoded" 'http://60.205.148.21:8080/post?id=100&page=1' -d "name=yaoxf&message=haha"
