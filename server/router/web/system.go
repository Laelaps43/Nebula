package web

import (
	"github.com/gin-gonic/gin"
	"nebula.xyz/api/v1/web"
)

type SystemRouter struct{}

func (s *SystemRouter) InitSystemRouter(Router *gin.RouterGroup) (r gin.IRouter) {
	systemRouter := Router.Group("system")
	systemApi := web.WebApiAll.SystemApi
	{
		systemRouter.GET("info", systemApi.GetSystemInfo)
	}
	return
}
