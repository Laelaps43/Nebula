package web

type webRouter struct {
	Hello
	UserRouter
	DeviceRouter
	ZlmHookRouter
	ChannelRouter
}

var WebRouterAll = new(webRouter)
