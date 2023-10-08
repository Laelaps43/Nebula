package web

type WebService struct {
	UserService
	JwtService
	DeviceService
}

var WebServiceAll = new(WebService)
