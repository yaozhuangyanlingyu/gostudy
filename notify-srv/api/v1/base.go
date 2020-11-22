package v1

import (
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
)

type Base struct {
	redisGo *redis.Client
	dbGo    *gorm.DB
}
