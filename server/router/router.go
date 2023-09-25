package router

import (
	"github.com/gin-gonic/gin"
	"nebula.xyz/global"
	"nebula.xyz/router/web"
)

// 初始化总路由
func Routers() *gin.Engine {
	// 创建默认路由
	Router := gin.Default()

	// 路由汇总
	webRouter := web.WebRouterAll

	// 基本路由，不用被鉴权
	PublicGroup := Router.Group(global.CONFIG.SERVER.RouterPrefix)
	{
		webRouter.InitHelloRouter(PublicGroup)
		webRouter.InitUserRouter(PublicGroup)

	}

	// PrivateGroup := Router.Group(global.CONFING.SERVER.RouterPrefix).Use().Use()
	// {

	// }

	global.Logger.Info("路由初始化成功！")
	return Router
}
