package web

type webRouter struct {
	Hello
	UserRouter
	DeviceRouter
}

var WebRouterAll = new(webRouter)
