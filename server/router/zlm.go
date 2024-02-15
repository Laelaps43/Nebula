package router

import (
	"github.com/gin-gonic/gin"
	"nebula.xyz/api/v1/zlm"
)

type ZlmHookRouter struct{}

func (z *ZlmHookRouter) InitZlmHookRouter(Router *gin.RouterGroup) (r gin.IRouter) {
	hookRouter := Router.Group("/index/hook")
	hookApi := zlm.ZLMApiAll.ZlmHookApi
	{
		hookRouter.POST("on_server_keepalive", hookApi.OnServerKeepalive)
		hookRouter.POST("on_publish", hookApi.OnPublish)
		hookRouter.POST("on_stream_changed", hookApi.OnStreamChanged)
		hookRouter.POST("on_record_mp4", hookApi.OnRecordMp4)
		hookRouter.POST("on_stream_not_found", hookApi.OnStreamNotFound)
	}
	return
}
