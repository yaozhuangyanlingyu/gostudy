package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 使用 SecureJSON 可以防止 JSON 劫持：
func main() {
	r := gin.Default()

	// You can also use your own secure json prefix
	// r.SecureJsonPrefix(")]}',\n")

	r.GET("/someJSON", func(c *gin.Context) {
		names := []string{"lena", "austin", "foo"}

		// Will output  :   while(1);["lena","austin","foo"]
		c.SecureJSON(http.StatusOK, names)
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}

// 如果给定结构是数组的话，默认会添加 "while(1)," 前缀到响应实体：
