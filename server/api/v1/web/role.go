package web

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"nebula.xyz/global"
	"nebula.xyz/model"
	"nebula.xyz/model/request"
	resp "nebula.xyz/model/response"
	"strconv"
)

type RoleApi struct{}

// CreateRole 创建角色
func (r RoleApi) CreateRole(ctx *gin.Context) {
	var createRole request.CreateRole

	err := ctx.BindJSON(&createRole)
	if err != nil {
		model.ErrorWithMessage("参数错误", ctx)
		return
	}
	err = roleService.CreateRole(createRole)
	if err != nil {
		model.ErrorWithMessage("设备", ctx)
		global.Logger.Error("创建角色失败", zap.Error(err))
		return
	}
	model.OKWithMessage("创建成功", ctx)
}

// GetRolePagination 分页获取的角色
func (r RoleApi) GetRolePagination(ctx *gin.Context) {
	var pagination request.Pagination

	err := ctx.ShouldBindJSON(&pagination)
	if err != nil {
		model.ErrorWithMessage("请检查分页数据", ctx)
		return
	}

	roles, total, err := roleService.GetRolePagination(pagination)
	if err != nil {
		global.Logger.Error("获取失败", zap.Error(err))
		model.ErrorWithMessage("获取角色列表失败", ctx)
		return
	}
	model.OkWithDetailed(resp.PaginationResult{
		List:  roles,
		Total: total,
	}, "获取成功", ctx)
}

// UpdateRole 更新角色
func (r RoleApi) UpdateRole(ctx *gin.Context) {

	var role request.UpdateRole
	err := ctx.ShouldBindJSON(&role)
	if err != nil {
		model.ErrorWithMessage(err.Error(), ctx)
		return
	}

	err = roleService.UpdateRole(role)

	if err != nil {
		global.Logger.Error("更新失败", zap.Error(err))
		model.ErrorWithMessage("更新失败", ctx)
	}
	model.OKWithMessage("更新成功", ctx)
}

// DeleteRole 删除角色
func (r RoleApi) DeleteRole(ctx *gin.Context) {

	roleId := ctx.Param("roleId")

	roleIdUint, err := strconv.ParseUint(roleId, 10, 64)
	if err != nil {
		model.ErrorWithMessage("Invalid role ID", ctx)
		return
	}

	err = roleService.DeleteRole(uint(roleIdUint))
	if err != nil {
		model.ErrorWithMessage("删除角色"+err.Error(), ctx)
		return
	}

	model.OKWithMessage("删除成功", ctx)
}
