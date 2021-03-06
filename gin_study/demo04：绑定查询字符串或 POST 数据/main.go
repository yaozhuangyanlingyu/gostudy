package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name     string    `form:"name"`
	Address  string    `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

func startPage(c *gin.Context) {
	var person Person
	// If `GET`, only `Form` binding engine (`query`) used.
	// If `POST`, first checks the `content-type` for `JSON` or `XML`, then uses `Form` (`form-data`).
	// See more at https://github.com/gin-gonic/gin/blob/master/binding/binding.go#L48
	err := c.ShouldBind(&person)
	if err == nil {
		log.Println(person.Name)
		log.Println(person.Address)
		log.Println(person.Birthday)
	} else {
		fmt.Println(err)
	}
	c.String(200, "Success")
}

func main() {
	route := gin.Default()
	// GET 请求
	route.GET("/testing", startPage)
	// POST 请求
	route.POST("/testing", startPage)
	route.Run(":8080")
}

// https://laravelacademy.org/post/21868
// http://app.dev02.aplum-inc.com:8080/testing?name=%E6%9C%B1%E5%BD%A9%E5%A8%9F&address=%E5%8C%97%E8%90%A5&birthday=2021-03-06
