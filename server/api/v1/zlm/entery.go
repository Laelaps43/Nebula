package zlm

import (
	"nebula.xyz/service"
)

// WebApi web全部Api对象
type zlmApi struct {
	ZlmHookApi
}

var (
	zlmService = service.SysServiceAll.ZLMService
)

var ZLMApiAll = new(zlmApi)
