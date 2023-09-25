package web

import "nebula.xyz/service/web"

// web全部Api对象
type WebApi struct {
	Hello
	UserApi
}

var (
	userService = web.WebServiceAll.UserService
	jwtService  = web.WebServiceAll.JwtService
)

var WebApiAll = new(WebApi)
