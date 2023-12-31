package service

type sysService struct {
	ZLMService
	CasbinService
}

var SysServiceAll = new(sysService)
