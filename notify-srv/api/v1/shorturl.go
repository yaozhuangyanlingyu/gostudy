package v1

import (
	"notify/pkg/logger"
	"notify/service"
	handleService "notify/service/shorturl"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
)

type ShortUrl struct {
	Base    *Base
	Service *handleService.ShortUrlService
}

func NewShortURL(redisGo *redis.Client, dbGo *gorm.DB) *ShortUrl {
	service := &handleService.ShortUrlService{
		Base: &service.Base{
			RedisGo: redisGo,
			DbGo:    dbGo,
		},
	}
	apiParams := &ShortUrl{
		Base: &Base{
			RedisGo: redisGo,
			DbGo:    dbGo,
		},
		Service: service,
	}
	return apiParams
}

func (this *ShortUrl) GenUrl(c *gin.Context) {
	longUrl := c.Query("long_url")
	passwd := c.Query("passwd")
	logger.Info("request shorturl", longUrl, passwd)
	if len(longUrl) == 0 {
		c.JSON(200, gin.H{
			"code": 500001,
			"msg":  "long_url参数不能为空",
			"data": "",
		})
		return
	}
	if passwd != "plum-url" {
		c.JSON(200, gin.H{
			"code": 500001,
			"msg":  "passwd参数错误",
			"data": "",
		})
		return
	}

	// 获取返回短链URL
	url, err := this.Service.Handle(longUrl)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 500001,
			"msg":  err.Error(),
			"data": "",
		})
		return
	}

	// 返回短链数据
	data := make(map[string]string)
	data["short_url"] = url
	data["long_url"] = longUrl
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": data,
	})
}
