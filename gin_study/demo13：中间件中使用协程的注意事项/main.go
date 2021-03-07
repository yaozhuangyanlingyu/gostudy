package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// 在中间件或处理器中开启新的协程时，不应该在其中使用原生的上下文对象，而应该使用它的只读副本

func main() {
	r := gin.Default()

	r.GET("/long_async", func(c *gin.Context) {
		// 在协程中使用上下文对象c的只读副本
		cCp := c.Copy()
		go func() {
			// 通过 time.Sleep() 模拟耗时任务
			time.Sleep(5 * time.Second)

			// ⚠️ 注意这里使用的是上下文对象的只读副本 "cCp"
			log.Println("Done! in path " + cCp.Request.URL.Path)
		}()
	})

	r.GET("/long_sync", func(c *gin.Context) {
		// 通过 time.Sleep() 模拟耗时任务
		time.Sleep(5 * time.Second)

		// 这里没有使用协程，所以可以直接使用上下文对象 c
		log.Println("Done! in path " + c.Request.URL.Path)
	})

	r.Run(":8080")
}
