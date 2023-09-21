package web

import (
	"github.com/gin-gonic/gin"
	web "nebula.xyz/api/v1/web"
)

type Hello struct{}

func (h *Hello)InitHelloRouter(Router *gin.RouterGroup) (R gin.IRouter){
	helloRouter := Router.Group("hello")
	helloApi := web.WebApiAll.Hello
	{
		helloRouter.GET("say", helloApi.SayHello)
	}
	return helloRouter
}