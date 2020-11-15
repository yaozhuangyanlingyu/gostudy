package api

import (
	"notify/api/shorturl"
	"notify/pkg/options"

	"github.com/gin-gonic/gin"
)

func RunGin(serverConf *options.GinServerConfig) {
	r := gin.Default()

	// 路由解析
	r.GET("/gen-url", shorturl.GenUrl)

	r.Run(serverConf.HTTPAddr)
}
