package main

import (
	"net/http"
	"time"
	"log"
	"github.com/gin-gonic/gin"
)

// StatCost 是一个统计耗时请求耗时的中间件
func StatCost() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Set("name", "小王子") // 可以通过c.Set在请求上下文中设置值，后续的处理函数能够取到该值
		// 调用该请求的剩余处理程序
		c.Next()
		// 不调用该请求的剩余处理程序
		// c.Abort()
		// 计算耗时
		cost := time.Since(start)
		log.Println(cost)
	}
}

func main() {
	r := gin.Default()

	// 注册一个全局中间件
	r.Use(StatCost())

	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/**/*")

	r.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.html", gin.H{
			"title": "posts/index",
		})
	})

	r.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.html", gin.H{
			"title": "users/index",
		})
	})

	r.GET("/test/json", func(c *gin.Context) {
		var msg struct {
			Name string `json:"user"`
			Message string
			Age int
		}
		msg.Name = "zhangfei"
		msg.Message = "Hello world"
		msg.Age = 30
		c.JSON(http.StatusOK, msg)
	})

	r.GET("/user/search", func(c *gin.Context){
		username := c.DefaultQuery("username", "小王子")
		address := c.Query("address")
		c.JSON(http.StatusOK, gin.H{
			"message" : "ok",
			"username" : username,
			"address" : address,
		})
	})

	r.POST("/user/search2", func(c *gin.Context) {
		// DefaultPostForm取不到值时会返回指定的默认值
		//username := c.DefaultPostForm("username", "小王子")
		username := c.PostForm("username")
		address := c.PostForm("address")
		//输出json结果给调用方
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"address":  address,
		})
	})

	r.GET("/user/search/:username/:address", func(c *gin.Context) {
		username := c.Param("username")
		address := c.Param("address")
		//输出json结果给调用方
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"address":  address,
		})
	})

	r.GET("/test", func(c *gin.Context) {
		name := c.MustGet("name").(string) // 从上下文取值
		log.Println(name)
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world!",
		})
	})

	r.Run()
}
