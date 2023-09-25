package jwt

import "github.com/golang-jwt/jwt/v4"

// 自定义声明，加入JWT载荷当中
type CustomClaims struct {
	ID       int
	UserName string
	jwt.RegisteredClaims
}
