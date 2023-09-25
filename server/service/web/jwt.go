package web

import (
	"nebula.xyz/global"
	"time"
)

// Jwt service

type JwtService struct{}

const JWT_PREFIX = "jwt:"

// 返回存在的Token，如果不存在返回

func (s JwtService) GetJWT(name string) (string, error) {
	result, err := global.CACHE.Get(JWT_PREFIX + name)
	return result.(string), err
}

func (s JwtService) SetJWT(token string, name string, expire time.Duration) (any, error) {
	return global.CACHE.Set(JWT_PREFIX+name, token, expire)
}

func (s JwtService) DeleteByKey(key string) error {
	return global.CACHE.DeleteByKey(key)
}
