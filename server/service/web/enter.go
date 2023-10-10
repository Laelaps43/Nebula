package web

type WebService struct {
	UserService
	JwtService
	DeviceService
	ZLMService
	VideoService
	ChannelService
}

var WebServiceAll = new(WebService)
