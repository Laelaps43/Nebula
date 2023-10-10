package router

import (
	"github.com/gin-gonic/gin"
	"nebula.xyz/global"
	"nebula.xyz/router/web"
)

// Routers 初始化总路由
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
		webRouter.InitVideoRouter(PublicGroup)
		webRouter.InitDeviceRouter(PublicGroup)
		webRouter.InitZlmHookRouter(PublicGroup)
	}

	// 鉴定路由
	//PrivateGroup := Router.Group(global.CONFIG.SERVER.RouterPrefix).Use(middleware.JWTAuth()).Use()
	//{

	//}

	global.Logger.Info("路由初始化成功！")
	return Router
}
