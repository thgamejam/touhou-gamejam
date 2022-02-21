package data

import (
	"github.com/go-redis/redis/v8"
	"service/pkg/conf"
)

// NewCache 初始化缓存
func NewCache(c *conf.Service) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Network:      c.Data.Redis.Network,
		Addr:         c.Data.Redis.Addr,
		Password:     c.Data.Redis.Password, // no password set
		ReadTimeout:  c.Data.Redis.ReadTimeout.AsDuration(),
		WriteTimeout: c.Data.Redis.WriteTimeout.AsDuration(),
		DB:           0, // use default DB
	})
	return rdb, nil
}
