package web

import "nebula.xyz/service/web"

// web全部Api对象
type WebApi struct {
	Hello
	UserApi
	VideoApi
	DeviceApi
}

var (
	userService   = web.WebServiceAll.UserService
	jwtService    = web.WebServiceAll.JwtService
	deviceService = web.WebServiceAll.DeviceService
)

var WebApiAll = new(WebApi)
