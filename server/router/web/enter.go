package web

type webRouter struct {
	Hello
	UserRouter
	RoleRouter
	HomeRouter
	DeviceRouter
	RecordRouter
	ChannelRouter
}

var WebRouterAll = new(webRouter)
