package utils

import (
	"github.com/gin-gonic/gin"
	"nebula.xyz/global"
	"nebula.xyz/model/jwt"
)

// GetClaims 从请求头中获取Claims
func GetClaims(c *gin.Context) (*jwt.CustomClaims, error) {
	token := c.Request.Header.Get("Authorization")
	j := NewJWT()
	claims, err := j.ParseToken(token)
	if err != nil {
		global.Logger.Error("未从请求头中获取到token，请检查Authorization是否携带token")
	}
	return claims, err
}
