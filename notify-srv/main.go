package main

import (
	"notify/pkg/config"

	"notify/pkg/logger"
	"notify/pkg/options"

	"github.com/gin-gonic/gin"
)

func init() {
	v, err := config.SetupConfig("./configs")
	if err != nil {
		panic(err)
	}
	config.Decode(v.AllSettings(), &options.Opts)
	logger.SetupLogger(&options.Opts.Log)
}

func main() {
	logger.Info("starting service...")

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(options.Opts.Server.HTTPAddr)
}
