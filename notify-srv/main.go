package main

import (
	"fmt"
	"notify/tools/config"

	"notify/tools/options"

	"github.com/gin-gonic/gin"
)

func init() {
	v, err := config.SetupConfig("./configs")
	if err != nil {
		panic(err)
	}
	config.Decode(v.AllSettings(), &options.Opts)

	fmt.Println(v.AllSettings())
	fmt.Println("====", options.Opts.Server.HTTPAddr, "====")
	fmt.Println("====", options.Opts.Server.Mode, "====")
	fmt.Println(options.Opts)
	fmt.Println("===========================================")
	//logger.InitLogger(&options.Opts.Log)
}

func main() {
	//fmt.Println(options.Opts)

	//logger.Info("starting service...")

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(options.Opts.Server.HTTPAddr)
}
