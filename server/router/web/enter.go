package web

type webRouter struct {
	Hello
	UserRouter
	RoleRouter
	HomeRouter
	DeviceRouter
	SystemRouter
	RecordRouter
	ChannelRouter
}

var WebRouterAll = new(webRouter)
