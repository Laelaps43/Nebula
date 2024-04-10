package web

import (
	"github.com/gin-gonic/gin"
	"nebula.xyz/model"
	"nebula.xyz/model/response"
	"nebula.xyz/utils"
)

func (u *UserApi) GetUserPermission(c *gin.Context) {

	claims, err := utils.GetClaims(c)
	if err != nil {
		model.ErrorWithMessage("获取用户信息失败", c)
		return
	}
	auths, modules, err := permissionService.GetAuths(claims.RoleId, claims.ID)
	if err != nil {
		model.ErrorWithMessage(err.Error(), c)
		return
	}

	model.OkWithDetailed(response.Permission{
		Modules: modules,
		Auths:   auths,
	}, "获取权限成功", c)
}
