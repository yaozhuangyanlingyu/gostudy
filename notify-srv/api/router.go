package api

import (
	"fmt"
	v1 "notify/api/v1"
	"notify/pkg/config"
	"notify/pkg/dbgo"
	"notify/pkg/goredis"
	"notify/pkg/logger"
	"notify/pkg/options"

	"github.com/gin-gonic/gin"
)

type Engine struct {
	shortUrlObj *v1.ShortUrl
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
	redisGo, err := goredis.NewRedis(&options.Opts.Redis)
	if err != nil {
		errMsg := "init redisClient err: " + err.Error()
		logger.Error(errMsg)
		panic(errMsg)
	}
	fmt.Println(redisGo)

	// 初始化mysql
	dbGo, err := dbgo.NewMysqlDB(&options.Opts.MySQL)
	if err != nil {
		errMsg := "init mysqlClient err: " + err.Error()
		logger.Error(err.Error())
		panic(errMsg)
	}
	fmt.Println(dbGo)

	// 实例化
	shortUrlObj := v1.NewShortUrl(redisGo, dbGo)
	this.shortUrlObj = shortUrlObj
}

func (this *Engine) RunEngine() {
	router := gin.Default()
	router.GET("/gen-url", this.shortUrlObj.GenUrl)
	router.Run()
}
