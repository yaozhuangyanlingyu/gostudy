package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 使用 JSONP 可以通过另一个域名从服务器请求数据，对应的 Gin 服务端代码可以这么写：
func main() {
	r := gin.Default()

	r.GET("/jsonp", func(c *gin.Context) {
		data := map[string]interface{}{
			"name": "学院君",
		}

		//如果传递的查询字符串 callback=hello
		// curl http://60.205.148.21:8080/jsonp?callback=hello
		//则输出：hello({\"name\":\"学院君\"})
		c.JSONP(http.StatusOK, data)
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
