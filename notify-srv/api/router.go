package api

import (
	apiV1 "notify/api/v1"
	"notify/pkg/config"
	"notify/pkg/dbgo"
	"notify/pkg/goredis"
	"notify/pkg/logger"
	"notify/pkg/options"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
)

type Engine struct {
	redisGo *redis.Client
	dbGo    *gorm.DB
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
	this.redisGo = redisGo

	// 初始化mysql
	dbGo, err := dbgo.NewMysqlDB(&options.Opts.MySQL)
	if err != nil {
		errMsg := "init mysqlClient err: " + err.Error()
		logger.Error(err.Error())
		panic(errMsg)
	}
	this.dbGo = dbGo
}

func (this *Engine) RunEngine() {
	router := gin.Default()

	// 短链处理
	shortUrlObj := apiV1.NewShortURL(this.redisGo, this.dbGo)
	router.GET("/gen-url", shortUrlObj.GenUrl)

	router.Run()
}
