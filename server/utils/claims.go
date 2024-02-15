package utils

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"nebula.xyz/global"
	"nebula.xyz/model/jwt"
	"nebula.xyz/model/system"
)

// GetClaims 从请求头中获取Claims
func GetClaims(c *gin.Context) (*jwt.CustomClaims, error) {
	token := c.Request.Header.Get("Authorization")
	j := NewJWT()
	claims, err := j.ParseToken(token)
	if err != nil {
		global.Logger.Error("未从请求头中获取到token，请检查Authorization是否携带token", zap.Error(err))
	}
	return claims, err
}

// CheckUserRole 检查用户是否拥有指定角色
func CheckUserRole(user *system.SysUser, roleID uint) bool {
	for _, r := range user.Roles {
		if r.ID == roleID {
			return true
		}
	}
	return false
}
