package v1

import (
	sShortUrl "notify/service/shorturl"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
)

type ShortUrl struct {
	base *Base
}

func NewShortUrl(redisGo *redis.Client, dbGo *gorm.DB) *ShortUrl {
	shortUrlObj := &ShortUrl{
		&Base{
			redisGo: redisGo,
			dbGo:    dbGo,
		},
	}
	return shortUrlObj
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
	url := sShortUrl.Handle(longUrl)
	data := make(map[string]string)
	data["url"] = url
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": data,
	})
}
