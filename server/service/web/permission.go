package web

import (
	"go.uber.org/zap"
	"nebula.xyz/global"
	"nebula.xyz/model/system"
)

type PermissionService struct{}

func (p *PermissionService) GetAuths() (auths []string, err error) {
	result := global.DB.Model(&system.SysButton{}).Pluck("slug", &auths)
	if result.Error != nil {
		global.Logger.Error("获取按钮权限失败", zap.Error(result.Error))
		return nil, result.Error
	}

	return auths, nil
}
