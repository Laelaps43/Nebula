package web

import "nebula.xyz/service/web"

// web全部Api对象
type WebApi struct {
	Hello
	UserApi
	VideoApi
	DeviceApi
	ZlmHookApi
	ChannelApi
}

var (
	userService    = web.WebServiceAll.UserService
	jwtService     = web.WebServiceAll.JwtService
	deviceService  = web.WebServiceAll.DeviceService
	zlmService     = web.WebServiceAll.ZLMService
	videoService   = web.WebServiceAll.VideoService
	channelService = web.WebServiceAll.ChannelService
)

var WebApiAll = new(WebApi)
