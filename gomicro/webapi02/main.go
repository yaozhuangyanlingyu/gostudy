package main

import(
	"github.com/micro/go-micro/web"
	"net/http"
	"github.com/gin-gonic/gin"
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
