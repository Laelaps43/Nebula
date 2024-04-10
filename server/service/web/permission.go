package web

import (
	"errors"
	"nebula.xyz/global"
	"nebula.xyz/model/system"
)

type PermissionService struct{}

func (p *PermissionService) GetAuths(roleId uint, userId uint) ([]string, []string, error) {
	var count int64
	err := global.DB.Model(&system.SysUserRole{}).
		Where("sys_user_id = ? AND sys_role_id = ?", userId, roleId).Count(&count).Error
	if err != nil {
		return nil, nil, errors.New("获取用户权限失败")
	}
	if count == 0 {
		return nil, nil, errors.New("获取用户权限失败")
	}
	var role system.SysRole
	err = global.DB.Model(&system.SysRole{}).
		Where("id = ?", roleId).
		Preload("SysMenus").
		Preload("SysButton").
		First(&role).Error
	if err != nil {
		return nil, nil, errors.New("获取用户权限失败")
	}
	modules := make([]string, 0)
	for _, role := range role.SysMenus {
		modules = append(modules, role.Slug)
	}
	auths := make([]string, 0)
	for _, button := range role.SysButton {
		auths = append(auths, button.Slug)
	}
	return auths, modules, nil
}
