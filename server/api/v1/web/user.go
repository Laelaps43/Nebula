package web

import (
	"github.com/gin-gonic/gin"
	jwt4 "github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
	"nebula.xyz/global"
	"nebula.xyz/model"
	"nebula.xyz/model/request"
	resp "nebula.xyz/model/response"
	"nebula.xyz/utils"
	"strconv"
	"time"
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

func (u *UserApi) GetUsrRole(ctx *gin.Context) {
	claims, err := utils.GetClaims(ctx)
	if err != nil {
		model.ErrorWithMessage("获取角色错误", ctx)
		return
	}
	role, err := userService.GetUserRole(claims.RoleId, claims.ID)
	if err != nil {
		model.ErrorWithMessage("获取角色错误", ctx)
		return
	}
	model.OkWithDetailed(role, "获取成功", ctx)
}

func (u *UserApi) SwitchRole(ctx *gin.Context) {
	var switchRole request.SwitchRole
	err := ctx.ShouldBindJSON(&switchRole)
	if err != nil {
		model.ErrorWithMessage("切换角色失败", ctx)
		return
	}
	claims, err := utils.GetClaims(ctx)
	token := ctx.Request.Header.Get("Authorization")
	err = userService.SwitchRole(switchRole.RoleId, claims.ID)
	if err != nil {
		model.ErrorWithMessage(err.Error(), ctx)
		return
	}
	claims.RoleId = switchRole.RoleId
	j := utils.NewJWT()
	dr, _ := utils.ParseExpireTime(global.CONFIG.JWT.JwtExpire)
	claims.ExpiresAt = jwt4.NewNumericDate(time.Now().Add(dr))
	newToken, _ := j.CreateTokenByOlderToken(token, claims)
	ctx.Header("new-token", newToken)
	ctx.Header("new-expire-at", strconv.FormatInt(claims.ExpiresAt.Unix(), 10))
	// 将新Token保存到Cache中，替换以前的Token
	if _, err := jwtService.SetJWT(token, strconv.Itoa(int(claims.ID)), dr); err != nil {
		global.Logger.Error("保存Token失败！", zap.Error(err))
		model.ErrorWithMessage("登录失败，请稍候再试！", ctx)
		return
	}
	model.OKWithMessage("切换成功", ctx)
}
