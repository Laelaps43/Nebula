package web

type WebService struct {
	UserService
	JwtService
}

var WebServiceAll = new(WebService)
