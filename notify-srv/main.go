package main

import (
	"notify/tools/config"

	"notify/tools/options"

	"notify/tools/logger"

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
	logger.Info("starting service...")
	defer logger.Sync()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(options.Opts.Server.HTTPAddr)
}
