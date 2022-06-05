package cache

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"service/pkg/conf"
	"time"
)

type Cache struct {
	Client *redis.Client // redis客户端

	DelayDeleteTime time.Duration // 延迟删除等待时间
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
		Client:          rdb,
		DelayDeleteTime: c.Data.Redis.DelayDeleteTime.AsDuration(),
	}, nil
}

// Set 保存数据
func (c *Cache) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	jsonByte, err := json.Marshal(value)
	if err != nil {
		return err
	}
	jsonString := string(jsonByte)
	return c.Client.Set(ctx, key, jsonString, expiration).Err()
}

func (c *Cache) SetString(ctx context.Context, key string, value string, expiration time.Duration) error {
	return c.Client.Set(ctx, key, value, expiration).Err()
}

// Get 获取数据
// value 必须是一个指针
func (c *Cache) Get(ctx context.Context, key string, value interface{}) (ok bool, err error) {
	v, err := c.Client.Get(ctx, key).Result()
	if err == redis.Nil {
		return false, nil
	} else if err != nil {
		return
	}
	return true, json.Unmarshal([]byte(v), value)
}

func (c *Cache) GetString(ctx context.Context, key string) (value string, ok bool, err error) {
	value, err = c.Client.Get(ctx, key).Result()
	if err == redis.Nil {
		return value, false, nil
	} else if err != nil {
		return value, false, err
	}
	return value, true, nil
}

// Del 删除数据
func (c *Cache) Del(ctx context.Context, key string) (err error) {
	_, err = c.Client.Del(ctx, key).Result()

	// TODO 简单的延迟删除
	var delayDel = func(key string) {
		time.Sleep(c.DelayDeleteTime)
		_, _ = c.Client.Del(context.Background(), key).Result()
	}
	go delayDel(key)

	return err
}

func (c *Cache) Close() error {
	return c.Client.Close()
}

func (c *Cache) GetClient() *redis.Client {
	return c.Client
}
