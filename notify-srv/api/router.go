package api

import (
	"fmt"
	v1 "notify/api/v1"
	"notify/pkg/config"
	"notify/pkg/goredis"
	"notify/pkg/logger"
	"notify/pkg/options"

	"github.com/gin-gonic/gin"
)

type Engine struct {
	genShortUrl *v1.GenShortUrl
}

func (this *Engine) InitRouter() {
	// 初始化配置
	v, err := config.SetupConfig("./configs")
	if err != nil {
		panic(err)
	}
	config.Decode(v.AllSettings(), &options.Opts)

	// 初始化日志库
	logger.SetupLogger(&options.Opts.Log)

	// 初始化redis
	redisClient, err := goredis.NewRedis(&options.Opts.Redis)
	if err != nil {
		errMsg := "init redisClient err: " + err.Error()
		logger.Error(errMsg)
		panic(errMsg)
	}
	fmt.Println(redisClient)

	// 初始化mysql

	// 实例化
}

func Hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 500001,
		"msg":  "long_url参数不能为空",
		"data": "",
	})
}

func (this *Engine) RunEngine() {
	router := gin.Default()
	router.GET("/hello", Hello)
	router.Run()
}
