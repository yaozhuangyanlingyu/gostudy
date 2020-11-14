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
	logger.SetupLogger()
}

func main() {
	//logger.Info("starting service... %s xxx", "党静角好想操你")
	//logger.Warn("starting service %s ... ggg", "党静角好想操你")
	//logger.Error("starting %s service...", "党静角好想操你")

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(options.Opts.Server.HTTPAddr)
}
