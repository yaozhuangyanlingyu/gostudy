package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// This handler will match /user/xueyuanjun but will not match /user/ or /user
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// However, this one will match /user/xueyuanjun/ and also /user/xueyuanjun/send
	// If no other routers match /user/xueyuanjun, it will redirect to /user/xueyuanjun/
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	router.Run(":8080")
}

// 运行上述代码，在终端窗口通过 curl 发起模拟请求进行测试：
// curl http://60.205.148.21:8080/user/guanyu
// curl http://60.205.148.21:8080/user/guanyu/send
