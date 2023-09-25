package pkg

import (
	"context"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"nebula.xyz/global"
	"strings"
)

func GetCache() global.Cache {
	CacheType := global.CONFIG.SERVER.CacheType
	// 系统指定缓存为Redis
	if strings.ToLower(CacheType) == "redis" {
		// 获取全局配置
		rediscfg := global.CONFIG.REDIS
		client := redis.NewClient(&redis.Options{
			Addr:     rediscfg.Addr,
			Password: rediscfg.Password,
			DB:       rediscfg.DB,
		})
		// 测试
		_, err := client.Ping(context.Background()).Result()

		if err != nil {
			global.Logger.Error("Failed to connect redis, err: ", zap.String("err", err.Error()))
		} else {
			global.Logger.Info("Cache Is Redis")
		}
		return NewRedisClient(client)
	} else {
		// 默认缓存
		return nil
	}
}
