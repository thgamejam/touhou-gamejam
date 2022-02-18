package cache

import (
	"github.com/go-redis/redis/v8"
	"service/pkg/conf"
)

// NewCache 初始化缓存
func NewCache(c *conf.Data_Redis) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Network:      c.Network,
		Addr:         c.Addr,
		Password:     c.Password, // no password set
		ReadTimeout:  c.ReadTimeout.AsDuration(),
		WriteTimeout: c.WriteTimeout.AsDuration(),
		DB:           0, // use default DB
	})
	return rdb, nil
}
