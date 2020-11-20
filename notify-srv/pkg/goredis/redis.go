package goredis

import (
	"time"

	"github.com/go-redis/redis/v8"
)

type Config struct {
	Host     string
	Password string
	Db       int
}

// NewRedis Initialize the Redis instance
func NewRedis(redisConf *Config) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:         redisConf.Host,
		Password:     redisConf.Password,
		DB:           redisConf.Db,
		ReadTimeout:  time.Minute,
		WriteTimeout: time.Minute,
	})

	if err := client.Ping(client.Context()).Err(); err != nil {
		return nil, err
	}

	return client, nil
}
