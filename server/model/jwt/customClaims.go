package jwt

import "github.com/golang-jwt/jwt/v4"

// CustomClaims 自定义声明，加入JWT载荷当中
type CustomClaims struct {
	ID       uint
	UserName string
	RoleId   uint
	jwt.RegisteredClaims
}
