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

// CreateClaims 自定义声明，将被载入到载荷中
func (j *JWT) CreateClaims(user *system.SysUser, expire time.Time) jwt.CustomClaims {
	return jwt.CustomClaims{
		ID:       user.ID,
		UserName: user.UserName,
		RoleId:   user.RoleID,
		RegisteredClaims: jwt4.RegisteredClaims{
			Audience:  jwt4.ClaimStrings{"Nebula"},     // 受众
			NotBefore: jwt4.NewNumericDate(time.Now()), // 生效时间
			ExpiresAt: jwt4.NewNumericDate(expire),     // 过期时间
			Issuer:    global.CONFIG.JWT.Issuer,        // 签名发行者
		},
	}
}

// CreateToken 创建token
func (j *JWT) CreateToken(claims jwt.CustomClaims) (string, error) {
	token := jwt4.NewWithClaims(jwt4.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey) // 签名
}

// ParseToken 解析Token
func (j *JWT) ParseToken(tokenString string) (*jwt.CustomClaims, error) {
	// 解析Token到CustomClaims
	token, err := jwt4.ParseWithClaims(tokenString, &jwt.CustomClaims{}, func(token *jwt4.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if value, ok := err.(*jwt4.ValidationError); ok {
			if value.Errors&jwt4.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if value.Errors&jwt4.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if value.Errors&jwt4.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*jwt.CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	} else {
		return nil, TokenInvalid
	}
}

// CreateTokenByOlderToken 根据原来的Token生产新Token
func (j *JWT) CreateTokenByOlderToken(olderToken string, claims *jwt.CustomClaims) (string, error) {
	v, err, _ := global.SingleFlight.Do("jwt:"+olderToken, func() (interface{}, error) {
		return j.CreateToken(*claims)
	})
	return v.(string), err
}
