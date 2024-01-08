package web

type WebService struct {
	JwtService
	UserService
	HomeService
	RoleService
	RecordServer
	VideoService
	DeviceService
	ChannelService
	PermissionService
}

var WebServiceAll = new(WebService)
