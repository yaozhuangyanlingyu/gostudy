package v1

import (
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
	if len(longUrl) == 0 {
		c.JSON(200, gin.H{
			"code": 500001,
			"msg":  "long_url参数不能为空",
			"data": "",
		})
		return
	}

	// 获取返回短链URL
	url := this.Service.Handle(longUrl)
	data := make(map[string]string)
	data["url"] = url
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": data,
	})
}
