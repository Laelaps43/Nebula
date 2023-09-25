package web

type webRouter struct {
	Hello
	UserRouter
}

var WebRouterAll = new(webRouter)
