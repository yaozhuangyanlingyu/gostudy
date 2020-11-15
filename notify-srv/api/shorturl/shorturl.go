package shorturl

import (
	sShortUrl "notify/service/shorturl"

	"github.com/gin-gonic/gin"
)

func GenUrl(c *gin.Context) {
	longUrl := c.Query("long_url")
	if len(longUrl) == 0 {
		c.JSON(200, gin.H{
			"code": 500001,
			"msg":  "",
			"data": "",
		})
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
