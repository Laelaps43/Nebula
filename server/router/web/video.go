package web

import (
	"github.com/gin-gonic/gin"
	"nebula.xyz/api/v1/web"
)

type Video struct{}

func (h *Hello) InitVideoRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	helloRouter := Router.Group("video")
	videoApi := web.WebApiAll.VideoApi
	{
		helloRouter.POST("play", videoApi.PlayVideo)
		helloRouter.POST("record", videoApi.RecordVideo)
	}
	return helloRouter
}
