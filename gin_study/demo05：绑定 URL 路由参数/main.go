package main

import "github.com/gin-gonic/gin"

type Student struct {
	ID   string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

func main() {
	route := gin.Default()
	route.GET("/:name/:id", func(c *gin.Context) {
		var student Student
		// 将路由参数绑定到结构体中
		if err := c.ShouldBindUri(&student); err != nil {
			c.JSON(400, gin.H{"msg": err})
			return
		}
		c.JSON(200, gin.H{"name": student.Name, "uuid": student.ID})
	})
	route.Run(":8080")
}

// curl
// http://app.dev02.aplum-inc.com:8080/zhucaijuan/327a1828-7224-48c8-8845-e31447214009
