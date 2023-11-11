package web

import (
	"errors"
	"fmt"
	"nebula.xyz/global"
	"nebula.xyz/model/system"
	"nebula.xyz/utils"
)

type UserService struct{}

// 用户登录
func (u *UserService) Login(user *system.SysUser) (userHandler *system.SysUser, err error) {
	if global.DB == nil {
		global.Logger.Error("数据库未配置")
		return nil, fmt.Errorf("数据库未配置")
	}
	var userTmp system.SysUser
	//TODO 这里可能需要同时去加载权限表
	err = global.DB.Where("email = ?", user.Email).First(&userTmp).Error
	if err == nil {
		// 判断密码是否相等
		if ok := utils.BcryptCheck(user.PassWord, userTmp.PassWord); !ok {
			return nil, errors.New("密码错误")
		}
	}
	return &userTmp, err
}
