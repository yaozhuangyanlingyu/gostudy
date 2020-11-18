package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func StatCost() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		cost := time.Since(start)
		log.Println(cost)
	}
}

func main() {
	r := gin.Default()

	shopGroup := r.Group("/shop", StatCost())
	{
		shopGroup.GET("/test1", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"name": "zhangsan",
				"age":  30,
			})
		})

		shopGroup.GET("/test2", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"name": "lisi",
				"age":  40,
			})
		})
	}

	r.Run()
}
