package middleware

// 处理Token中间件

import (
	"errors"
	"github.com/gin-gonic/gin"
	jwt4 "github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
	"nebula.xyz/global"
	"nebula.xyz/model"
	"nebula.xyz/service/web"
	"nebula.xyz/utils"
	"strconv"
	"time"
)

var jwtService = web.WebServiceAll

// JWTAuth JWT处理
func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 从Http头部Authorization获取到token
		token := ctx.Request.Header.Get("Authorization")
		if token == "" {
			// token为空
			model.ErrorWithMessage("请登录！", ctx)
			ctx.Abort()
			return
		}
		// 判断token是否有效
		// TODO 是否需要考虑Token是否在黑名单内

		j := utils.NewJWT()
		claims, err := j.ParseToken(token)
		if err != nil {
			// token 错误
			if errors.Is(err, utils.TokenExpired) {
				// Token 已过期
				model.ErrorWithDetailed(gin.H{"reload": true}, "Token已到期", ctx)
				ctx.Abort()
				return
			}
			model.ErrorWithDetailed(gin.H{"reload": true}, err.Error(), ctx)
			ctx.Abort()
			return
		}
		if claims.ExpiresAt.Unix()-time.Now().Unix() < 10000 {
			// Token 有效，但小于最后有效时间
			dr, _ := utils.ParseExpireTime(global.CONFIG.JWT.JwtExpire)
			claims.ExpiresAt = jwt4.NewNumericDate(time.Now().Add(dr))
			newToken, _ := j.CreateTokenByOlderToken(token, claims)
			ctx.Header("new-token", newToken)
			ctx.Header("new-expire-at", strconv.FormatInt(claims.ExpiresAt.Unix(), 10))
			// 将新Token保存到Cache中，替换以来的Token

			if _, err := jwtService.SetJWT(token, strconv.Itoa(claims.ID), dr); err != nil {
				global.Logger.Error("保存Token失败！", zap.Error(err))
				model.ErrorWithMessage("登录失败，请稍候再试！", ctx)
				return
			}
		}
		// TODO
		ctx.Next()
	}
}
