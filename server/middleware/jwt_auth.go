package middleware
// 处理Token中间件

import (
	"github.com/gin-gonic/gin"
	"nebula.xyz/model"
)


func JWTAuth() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		// 从Http头部Authoration获取到token
		token := ctx.Request.Header.Get("Authoration")
		if token == ""{
			// token为空
			model.ErrorWithMessage("请登录！", ctx)
			ctx.Abort()
			return
		}
		// 判断token是否有效

		// TODO 
		ctx.Next()
	}	
}