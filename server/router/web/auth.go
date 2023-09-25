package web

import (
	"github.com/gin-gonic/gin"
	"nebula.xyz/api/v1/web"
)

type UserRouter struct{}

func (a *UserRouter) InitUserRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	userRouter := Router.Group("")
	userApi := web.WebApiAll.UserApi
	{
		userRouter.POST("login", userApi.DoLogin)
	}
	return userRouter
}
