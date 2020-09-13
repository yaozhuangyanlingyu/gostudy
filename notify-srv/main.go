package main

import (
	"notify/options"
	"notify/tools/config"

	"github.com/gin-gonic/gin"
)

func init() {
	v, err := config.SetupConfig("./configs")
	if err != nil {
		panic(err)
	}
	config.Decode(v.AllSettings(), &options.Opts)
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
