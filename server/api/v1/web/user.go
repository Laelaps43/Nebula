package web

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"nebula.xyz/global"
	"nebula.xyz/model"
	"nebula.xyz/model/request"
	resp "nebula.xyz/model/response"
	"nebula.xyz/utils"
)

// 用户相关API
// 用户登录

type UserApi struct{}

// GetUserInfoPagination 分页获取用户数据
func (u *UserApi) GetUserInfoPagination(ctx *gin.Context) {

	var pagination request.Pagination

	err := ctx.ShouldBindJSON(&pagination)
	if err != nil {
		model.ErrorWithMessage("请检查分页数据", ctx)
		return
	}
	users, total, err := userService.GetUserInfoPagination(pagination)
	if err != nil {
		global.Logger.Error("查询用户信息失败")
		model.ErrorWithMessage("获取失败", ctx)
		return
	}
	model.OkWithDetailed(resp.PaginationResult{
		List:  users,
		Total: total,
	}, "获取成功", ctx)
}

// CreateUser 创建用户 编辑用户
func (u *UserApi) CreateUser(ctx *gin.Context) {

	var userCreate request.UserCreate

	err := ctx.ShouldBindJSON(&userCreate)
	if err != nil {
		model.ErrorWithMessage("请检查创建数据", ctx)
		return
	}
	err = userService.CreateUser(userCreate)
	if err != nil {
		global.Logger.Error("创建用户失败", zap.Error(err))
		model.ErrorWithMessage("创建用户失败："+err.Error(), ctx)
		return
	}
	model.OKWithMessage("创建用户成功", ctx)
}

func (u *UserApi) EnableUser(ctx *gin.Context) {
	var enableUser request.EnableUser

	err := ctx.ShouldBindJSON(&enableUser)
	if err != nil {
		model.ErrorWithMessage("请检查数据", ctx)
		return
	}

	err = userService.EnableUser(enableUser)
	if err != nil {
		model.ErrorWithMessage(err.Error(), ctx)
		return
	}
	model.OKWithMessage("更新成功", ctx)
}

// UpdateUser 编辑用户
func (u *UserApi) UpdateUser(ctx *gin.Context) {

	var userCreate request.UserUpdate

	err := ctx.ShouldBindJSON(&userCreate)
	if err != nil {
		global.Logger.Error("绑定数据错误", zap.Error(err))
		model.ErrorWithMessage("请检查创建数据", ctx)
		return
	}
	err = userService.UpdateUser(userCreate)
	if err != nil {
		global.Logger.Error("更新用户失败:", zap.Error(err))
		model.ErrorWithMessage("更新用户失败："+err.Error(), ctx)
		return
	}
	model.OKWithMessage("更新用户成功", ctx)
}

// DeleteUser 删除用户
func (u *UserApi) DeleteUser(ctx *gin.Context) {

	var userId string

	userId = ctx.Param("userId")
	if len(userId) == 0 {
		model.ErrorWithMessage("用户Id不正确", ctx)
	}

	err := userService.DeleteUser(userId)
	if err != nil {
		global.Logger.Error("删除用户失败:", zap.Error(err))
		model.ErrorWithMessage("更新用户失败："+err.Error(), ctx)
		return
	}
	model.OKWithMessage("删除用户成功", ctx)
}

func (u *UserApi) GetLoginUserInfo(ctx *gin.Context) {
	// 获取到用户的信息
	claims, err := utils.GetClaims(ctx)
	if err != nil {
		global.Logger.Error("获取用户信息错误", zap.Error(err))
		model.ErrorWithMessage("获取用户信息错误", ctx)
		return
	}
	user, err := userService.GetUserInfo(claims.ID)

	if err != nil {
		model.ErrorWithMessage(err.Error(), ctx)
		return
	}
	model.OkWithDetailed(user, "查询成功", ctx)
}

func (u *UserApi) EditUser(ctx *gin.Context) {
	var editUser request.EditUser

	err := ctx.ShouldBindJSON(&editUser)
	if err != nil {
		global.Logger.Error("绑定参数错误", zap.Error(err))
		model.ErrorWithMessage("参数错误", ctx)
		return
	}

	err = userService.EditUser(editUser)
	if err != nil {
		model.ErrorWithMessage(err.Error(), ctx)
	}
	model.OKWithMessage("修改成功", ctx)
}
