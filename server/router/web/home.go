package web

import (
	"github.com/gin-gonic/gin"
	"nebula.xyz/api/v1/web"
)

type HomeRouter struct{}

func (h *HomeRouter) InitHomeRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	homeRouter := Router.Group("home")
	homeApi := web.WebApiAll.HomeApi
	{
		homeRouter.GET("overview/info", homeApi.GetOverView)
		homeRouter.GET("system/info", homeApi.GetSystemInfo)
		homeRouter.GET("server/info", homeApi.GetServerInfo)
	}
	return homeRouter
}
