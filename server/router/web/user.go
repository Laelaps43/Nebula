package web

import (
	"github.com/gin-gonic/gin"
	"nebula.xyz/api/v1/web"
)

type UserRouter struct{}

func (a *UserRouter) InitUserRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	userRouter := Router.Group("/user")
	userApi := web.WebApiAll.UserApi
	{
		userRouter.POST("login", userApi.DoLogin)
		userRouter.POST("page", userApi.GetUserInfoPagination)
		userRouter.POST("create", userApi.CreateUser)
		userRouter.POST("enable", userApi.EnableUser)
		userRouter.POST("update", userApi.UpdateUser)
		userRouter.GET("delete/:userId", userApi.DeleteUser)
		userRouter.GET("my/info", userApi.GetLoginUserInfo)
		userRouter.POST("my/edit", userApi.EditUser)
		userRouter.GET("role", userApi.GetUsrRole)
		userRouter.POST("switchRole", userApi.SwitchRole)
	}
	return userRouter
}

func (a *UserRouter) InitUserAuthorizationRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	userRouter := Router.Group("/user")
	userApi := web.WebApiAll.UserApi
	{
		userRouter.GET("permission", userApi.GetUserPermission)
	}
	return userRouter
}
