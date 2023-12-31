package web

import "nebula.xyz/service"

type WebService struct {
	JwtService
	service.ZLMService
	UserService
	HomeService
	RoleService
	VideoService
	DeviceService
	ChannelService
	PermissionService
}

var WebServiceAll = new(WebService)
