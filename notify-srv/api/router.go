/*
 * @Author: yaoxf
 * @Date: 2020-11-15 19:03:18
 * @LastEditTime: 2020-11-23 01:37:47
 * @LastEditors: Please set LastEditors
 * @Description: 路由控制
 * @FilePath: \notify-srv\api\router.go
 */
package api

import (
	"context"
	"net/http"
	apiV1 "notify/api/v1"
	"notify/pkg/config"
	"notify/pkg/dbgo"
	"notify/pkg/goredis"
	"notify/pkg/logger"
	"notify/pkg/options"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
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

// timeout middleware wraps the request context with a timeout
func timeoutMiddleware(timeout time.Duration) func(c *gin.Context) {
	return func(c *gin.Context) {

		// wrap the request context with a timeout
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)

		defer func() {
			// check if context timeout was reached
			if ctx.Err() == context.DeadlineExceeded {

				// write response and abort the request
				c.Writer.WriteHeader(http.StatusGatewayTimeout)
				c.Abort()
			}

			//cancel to clear resources after finished
			cancel()
		}()

		// replace request with context wrapped request
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func (this *Engine) RunEngine() {
	gin.SetMode(options.Opts.Server.Mode)
	router := gin.Default()
	router.Use(logger.GinLogger(), logger.GinRecovery(true), timeoutMiddleware(time.Second*2))

	// 短链处理
	shortUrlObj := apiV1.NewShortURL(this.redisGo, this.dbGo)
	router.GET("/gen-url", shortUrlObj.GenUrl)

	logger.Info("Start Gin Server listening", zap.String("addr", options.Opts.Server.HTTPAddr))
	router.Run(options.Opts.Server.HTTPAddr)
}
