package web

import (
	service "nebula.xyz/service/web"
)

// WebApi web全部Api对象
type webApi struct {
	Hello
	HomeApi
	UserApi
	RoleApi
	VideoApi
	RecordApi
	DeviceApi
	ChannelApi
}

var (
	userService       = service.WebServiceAll.UserService
	roleService       = service.WebServiceAll.RoleService
	jwtService        = service.WebServiceAll.JwtService
	homeService       = service.WebServiceAll.HomeService
	recordService     = service.WebServiceAll.RecordServer
	deviceService     = service.WebServiceAll.DeviceService
	videoService      = service.WebServiceAll.VideoService
	channelService    = service.WebServiceAll.ChannelService
	permissionService = service.WebServiceAll.PermissionService
)

var WebApiAll = new(webApi)
