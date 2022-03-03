package cache

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"service/pkg/conf"
	"time"
)

type Cache struct {
	client *redis.Client
}

// NewCache 初始化缓存
func NewCache(c *conf.Service) (*Cache, error) {
	rdb := redis.NewClient(&redis.Options{
		Network:      c.Data.Redis.Network,
		Addr:         c.Data.Redis.Addr,
		Password:     c.Data.Redis.Password, // no password set
		ReadTimeout:  c.Data.Redis.ReadTimeout.AsDuration(),
		WriteTimeout: c.Data.Redis.WriteTimeout.AsDuration(),
		DB:           0, // use default DB
	})
	return &Cache{
		client: rdb,
	}, nil
}

// Save 保存数据
func (c *Cache) Save(ctx context.Context, key string, data interface{}, expiration time.Duration) error {
	jsonByte, err := json.Marshal(data)
	if err != nil {
		return err
	}
	jsonString := string(jsonByte)
	return c.client.Set(ctx, key, jsonString, expiration).Err()
}

func (c *Cache) SaveString(ctx context.Context, key string, value string, expiration time.Duration) error {
	return c.client.Set(ctx, key, value, expiration).Err()
}

// Get 获取数据
func (c *Cache) Get(ctx context.Context, key string, v interface{}) (ok bool, err error) {
	value, err := c.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return false, nil
	} else if err != nil {
		return
	}
	return true, json.Unmarshal([]byte(value), v)
}

func (c *Cache) GetString(ctx context.Context, key string) (value string, ok bool, err error) {
	value, err = c.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return value, false, nil
	} else if err != nil {
		return value, false, err
	}
	return value, true, nil
}

func (c *Cache) Close() error {
	return c.client.Close()
}

func (c *Cache) GetClient() *redis.Client {
	return c.client
}
