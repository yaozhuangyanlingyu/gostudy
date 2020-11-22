package shorturl

import (
	"context"
	"notify/model"
	"notify/service"
	"time"
)

type ShortUrlService struct {
	Base *service.Base
}

func (this *ShortUrlService) Handle(longUrl string) (string, error) {
	// 判断之前是否生成过
	existsShortUrl := model.GetShortUrlByLongURL(this.Base.DbGo, longUrl)
	return existsShortUrl, nil

	// 获取随机数字
	randNum := this.getRandNum()

	// 生成短链码
	shortCode := this.getShortCode(randNum)

	// 存入数据库
	_, err := model.SaveShortUrlCode(this.Base.DbGo, shortCode, longUrl)
	if err != nil {
		return "", err
	}
	return shortCode, nil
}

/**
 * @desc 生成一个不重复的随机数
 */
func (this *ShortUrlService) getRandNum() int64 {
	var startNum int64 = 1570204800
	incrKey := "short_url_randnum"
	incrNum := this.Base.RedisGo.Incr(context.Background(), incrKey).Val()
	randNum := time.Now().Unix() + incrNum - startNum
	return randNum
}

/**
 * @desc 数字转换成字符串
 */
func (this *ShortUrlService) getShortCode(randNum int64) string {
	table := []byte{
		'H', 'G', 'm', 'U', 'R', '4', 'p', 'J', 'y', 'd', '9', 'o', 'P', 'X', 'Q', 'q',
		's', 'V', 'Y', 'F', 'A', '2', 'b', 'g', 'M', 'i', 'w', 'I', 'h', 'c', 'D', '3',
		'j', '8', 'T', '0', 'x', 'S', '7', 'l', 'W', 'C', 'E', 'f', '5', '6', 'a', 'K',
		't', 'L', 'e', 'O', 'Z', 'r', 'v', 'u', 'z', 'N', 'k', '1', 'B', 'n',
	}
	shortCode := ""
	var binary int64 = 62
	for randNum > binary {
		x := randNum % binary
		shortCode = string(table[x]) + shortCode
		randNum = randNum / binary
	}
	if randNum > 0 {
		shortCode = string(table[randNum]) + shortCode
	}
	return shortCode
}
