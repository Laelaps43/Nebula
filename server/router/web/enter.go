package web

type webRouter struct {
	Hello
	UserRouter
	DeviceRouter
	ZlmHookRouter
}

var WebRouterAll = new(webRouter)
