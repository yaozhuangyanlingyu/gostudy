/*
 * @Author: yaoxf
 * @Date: 2020-11-22 17:52:45
 * @LastEditTime: 2020-11-22 21:36:48
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \notify-srv\model\TShortUrl.go
 */

package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type TShortUrl struct {
	ID         int64
	Url        string
	Code       string
	CreateTime int64
}

func (TShortUrl) TableName() string {
	return "t_short_url"
}

/**
 * 保存短链code码
 * @return TShortUrl
 */
func SaveShortUrlCode(dbGo *gorm.DB, shortUrlCode string, longUrl string) (*TShortUrl, error) {
	shortUrl := &TShortUrl{
		Url:        longUrl,
		Code:       shortUrlCode,
		CreateTime: time.Now().Unix(),
	}
	err := dbGo.Create(&shortUrl).Error
	if err != nil {
		return nil, err
	}
	return shortUrl, nil
}

/**
 * 根据长url查询短连接
 * @return TShortUrl
 */
func GetShortUrlByLongURL(dbGo *gorm.DB, longUrl string) string {
	data := &TShortUrl{}
	err := dbGo.Where("url = ?", longUrl).First(data).Error
	if err != nil {
		return ""
	}
	return data.Code
}
