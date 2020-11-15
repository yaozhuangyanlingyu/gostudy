package shorturl

import (
	sShortUrl "notify/service/shorturl"

	"github.com/gin-gonic/gin"
)

func GenUrl(c *gin.Context) {
	// 获取短链URL
	url := sShortUrl.Handle()

	c.JSON(200, gin.H{
		"url": url,
	})
}
