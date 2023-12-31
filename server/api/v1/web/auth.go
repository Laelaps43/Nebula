package web

// 用户登录，退出

import (
	"fmt"
	"nebula.xyz/model/response"
	"nebula.xyz/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"nebula.xyz/global"
	"nebula.xyz/model"
	sysRequest "nebula.xyz/model/request"
	"nebula.xyz/model/system"
)

// DoLogin 登录API
func (u *UserApi) DoLogin(c *gin.Context) {
	var userRequest sysRequest.Login // 登录信息

	err := c.ShouldBindJSON(&userRequest) // 获取到前端传来的数据
	//clientIp := c.ClientIP()
	if err != nil {
		model.ErrorWithMessage("帐号错误或者密码错误！", c)
		return
	}
	global.Logger.Info("info", zap.String("email", userRequest.Email), zap.String("pass", userRequest.Password))

	//TODO 这里需要去判断穿过来的时间准确性，帐号密码不为空

	// 从设置回去最大次数，防止爆次数
	loginNum := global.CONFIG.SERVER.LoginMaxNum
	// TODO 规定时间登录次数
	// loginTimeout := global.CONFING.SERVER.LoginTimeout

	//TODO 判断是否在黑名单内

	//TODO 需要不需要去考虑规定时间内，最大的请求次数，或者是对某一个操作的请求次数

	//TODO 登录认证
	uRequest := &system.SysUser{Email: userRequest.Email, PassWord: userRequest.Password}
	user, err := userService.Login(uRequest)
	if err != nil {
		global.Logger.Error("登录失败，用户名或密码错误！", zap.Error(err))

		model.ErrorWithMessage("用户名或者密码错误！", c)
		return
	}
	if user.Enable != 1 {
		// 用户被禁止登录
		global.Logger.Info("用户被禁止登录！", zap.Uint("用户名Id", user.ID))
		model.ErrorWithMessage(fmt.Sprintf("登录失败，用户%s被禁止登录", user.UserName), c)
		return
	}

	// 每次登录一次加入一次，返回加入之后的登录次数
	//TODO 这里应该带个超时时间
	if maxNum, err := global.CACHE.Increment(strconv.Itoa(int(user.ID))); err != nil {
		global.Logger.Error("登录错误！", zap.String("Login Error", err.Error()))
		model.ServerError(c)
		return
	} else {
		if maxNum > int64(loginNum) {
			global.Logger.Info("登录超过最大次数", zap.Int("LoginMaxNum", int(loginNum)), zap.Int64("NowMaxLoginNum", maxNum))
			model.ErrorWithMessage("登录过多，请稍候再试！", c)
			return
		}
	}
	global.Logger.Info("验证成功，正在生成token.....")

	global.Logger.Error("user" + strconv.Itoa(int(user.ID)) + user.Email)
	u.TokenGen(c, user)

}

// TokenGen 生成Token
func (u *UserApi) TokenGen(c *gin.Context, user *system.SysUser) {
	j := utils.NewJWT()
	// 处理Token到期时间
	eq, _ := utils.ParseExpireTime(global.CONFIG.JWT.JwtExpire)
	claims := j.CreateClaims(user, time.Now().Add(eq))

	// 创建Token
	token, err := j.CreateToken(claims)
	if err != nil {
		global.Logger.Error("获取Token失败！", zap.Error(err))
		model.ErrorWithMessage("获取Token失败", c)
		return
	}
	if jwtStr, err := jwtService.GetJWT(strconv.Itoa(int(user.ID))); err == global.CACHENil {
		global.Logger.Info("key不存在cache当中，将token保存到cache中")
		if _, err := jwtService.SetJWT(token, strconv.Itoa(int(user.ID)), eq); err != nil {
			global.Logger.Error("保存Token失败！", zap.Error(err))
			model.ErrorWithMessage("登录失败，请稍候再试！", c)
			return
		}
	} else if err != nil {
		global.Logger.Error("保存Token失败！", zap.Error(err))
		model.ErrorWithMessage("登录失败，请稍候再试！", c)
		return
	} else {
		global.Logger.Info("key以及存在Cache中，将原先的删除，然后在添加新Token")
		if err := jwtService.DeleteByKey(jwtStr); err != nil {
			global.Logger.Error("删除已存在的Key失败！", zap.String("key", jwtStr))
			model.ServerError(c)
			return
		}
		if _, err := jwtService.SetJWT(token, strconv.Itoa(int(user.ID)), eq); err != nil {
			global.Logger.Error("保存Token失败！", zap.Error(err))
			model.ErrorWithMessage("登录失败，请稍候再试！", c)
			return
		}
	}
	model.OkWithDetailed(response.LoginResponse{
		Token:     token,
		ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
	}, "登录成功", c)
	global.Logger.Info("登录成功", zap.String("email", user.Email))
}
