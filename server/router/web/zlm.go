package web

import (
	"github.com/gin-gonic/gin"
	"nebula.xyz/api/v1/web"
)

type ZlmHookRouter struct{}

func (z *ZlmHookRouter) InitZlmHookRouter(Router *gin.RouterGroup) (r gin.IRouter) {
	hookRouter := Router.Group("index/hook")
	hookApi := web.WebApiAll.ZlmHookApi
	{
		hookRouter.POST("on_server_keepalive", hookApi.OnServerKeepalive)
	}
	return
}
