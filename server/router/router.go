package router

import (
	"github.com/gin-gonic/gin"
	"nebula.xyz/global"
	"nebula.xyz/middleware"
	"nebula.xyz/router/web"
)

// Routers 初始化总路由
func Routers() *gin.Engine {
	// 创建默认路由
	Router := gin.Default()

	// 路由汇总
	webRouter := web.WebRouterAll
	zlmRouter := ZlmHookRouter{}
	global.Logger.Error("routerPrefix" + global.CONFIG.SERVER.RouterPrefix)
	// 基本路由，不用被鉴权
	PublicGroup := Router.Group(global.CONFIG.SERVER.RouterPrefix)
	{
		webRouter.InitHelloRouter(PublicGroup)
		webRouter.InitRoleRouter(PublicGroup)
		webRouter.InitUserRouter(PublicGroup)
		webRouter.InitVideoRouter(PublicGroup)
		webRouter.InitDeviceRouter(PublicGroup)
		webRouter.InitChannelRoute(PublicGroup)
		webRouter.InitRecordRouter(PublicGroup)
	}
	// ZLM webhook路由
	ZLMediaKitGroup := Router.Group("")
	{
		zlmRouter.InitZlmHookRouter(ZLMediaKitGroup)
	}

	// 权限路由
	AuthorizationGroup := Router.Group(global.CONFIG.SERVER.RouterPrefix)
	AuthorizationGroup.Use(middleware.JWTAuth())
	{
		webRouter.InitUserAuthorizationRouter(AuthorizationGroup)
	}

	// 鉴定路由
	PrivateGroup := Router.Group(global.CONFIG.SERVER.RouterPrefix)
	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{

		webRouter.InitHomeRouter(PrivateGroup)
	}

	global.Logger.Info("路由初始化成功！")
	return Router
}
