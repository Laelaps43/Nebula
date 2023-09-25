package global

// 初始化缓存，根据配置来判断是什么存储引擎

import (
	"time"
)

type Cache interface {
	Get(key string) (interface{}, error)
	Set(key string, value any, expire time.Duration) (any, error)
	DeleteByKey(keys ...string) error
	Increment(key string) (int64, error)
}

// CacheError 自定义错误类型，表示缓存中不存在指定key的缓存
type CacheError string

func (c CacheError) Error() string {
	return string(c)
}

const CACHENil = CacheError("the key not is exist in redis")
