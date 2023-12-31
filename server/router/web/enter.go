package web

import "nebula.xyz/router"

type webRouter struct {
	Hello
	UserRouter
	RoleRouter
	HomeRouter
	DeviceRouter
	SystemRouter
	router.ZlmHookRouter
	ChannelRouter
}

var WebRouterAll = new(webRouter)
