package service

import (
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
)

type Base struct {
	RedisGo *redis.Client
	DbGo    *gorm.DB
}
