package pkg

import (
	"context"
	"nebula.xyz/global"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	Connect *redis.Client
}

const REDIS_PREFIX = "nebula:"

var redisCon RedisClient
var ctx context.Context

// NewRedisClient 创建一个新Redis 客户端
func NewRedisClient(db *redis.Client) global.Cache {
	ctx = context.Background()
	redisCon = RedisClient{Connect: db}
	return &redisCon
}

// DeleteByKey 删除指定的key
func (r *RedisClient) DeleteByKey(keys ...string) error {
	return redisCon.Connect.Del(ctx, keys...).Err()
}

// Get 根据key获取数据
func (r *RedisClient) Get(key string) (any, error) {
	var err error
	result := redisCon.Connect.Get(ctx, REDIS_PREFIX+key)
	if result.Err() == redis.Nil {
		// 对Redis.Nil进行封装
		err = global.CACHENil
	} else {
		err = result.Err()
	}
	return result.Val(), err
}

// Set 存储key-value，以及TTL
func (r *RedisClient) Set(key string, value any, expire time.Duration) (any, error) {
	status := redisCon.Connect.Set(ctx, REDIS_PREFIX+key, value, expire)
	return status.Val(), status.Err()
}

// Increment 对指定的Key加1
func (r *RedisClient) Increment(key string) (int64, error) {
	intCmd := redisCon.Connect.Incr(ctx, REDIS_PREFIX+key)
	return intCmd.Val(), intCmd.Err()
}
