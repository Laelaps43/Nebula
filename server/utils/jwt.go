package utils

import (
	"errors"
	jwt4 "github.com/golang-jwt/jwt/v4"
	"nebula.xyz/global"
	"nebula.xyz/model/jwt"
	"nebula.xyz/model/system"
	"time"
)

type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired     error = errors.New("Token is expired")
	TokenNotValidYet error = errors.New("Token not active yet")
	TokenMalformed   error = errors.New("That's not even a token")
	TokenInvalid     error = errors.New("Couldn't handle this token:")
)

// 创建一个JWT对象

func NewJWT() *JWT {
	return &JWT{
		[]byte(global.CONFIG.JWT.SigningKey),
	}
}

// 自定义声明，将被载入到载荷中

func (j *JWT) CreateClaims(user *system.User, expire time.Time) jwt.CustomClaims {
	return jwt.CustomClaims{
		ID:       user.ID,
		UserName: user.UserName,
		RegisteredClaims: jwt4.RegisteredClaims{
			Audience:  jwt4.ClaimStrings{"Nebula"},     // 受众
			NotBefore: jwt4.NewNumericDate(time.Now()), // 生效时间
			ExpiresAt: jwt4.NewNumericDate(expire),     // 过期时间
			Issuer:    global.CONFIG.JWT.Issuer,        // 签名发行者
		},
	}
}

// 创建token

func (j *JWT) CreateToken(claims jwt.CustomClaims) (string, error) {
	token := jwt4.NewWithClaims(jwt4.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey) // 签名
}
