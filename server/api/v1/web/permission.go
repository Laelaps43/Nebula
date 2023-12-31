package web

import (
	"github.com/gin-gonic/gin"
	"nebula.xyz/model"
	"nebula.xyz/model/response"
)

func (u *UserApi) GetUserPermission(c *gin.Context) {

	modules := make([]string, 0)

	auths, _ := permissionService.GetAuths()

	model.OkWithDetailed(response.Permission{
		Modules: modules,
		Auths:   auths,
	}, "获取权限成功", c)
}
