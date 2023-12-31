package web

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"nebula.xyz/global"
	"nebula.xyz/model/request"
	"nebula.xyz/model/system"
)

type RoleService struct{}

func (s *RoleService) CreateRole(role request.CreateRole) error {
	sysRole := system.SysRole{}
	sysRole.Slug = role.Slug
	// 判断标识是否唯一
	err := sysRole.GetRoleBySlug()
	if err != nil {
		return err
	}
	if sysRole.ID != 0 {
		return errors.New(fmt.Sprintf("%s标识已经存在", role.Slug))
	}
	sysRole.Name = role.Name
	sysRole.ParentId = role.ParentId
	sysRole.Desc = role.Desc
	err = sysRole.SaveRole()
	if err != nil {
		return err
	}
	return nil
}

func (s *RoleService) GetRolePagination(pagination request.Pagination) (roles []system.SysRole, total int64, err error) {
	db := global.DB.Model(&system.SysRole{})

	err = db.Where("parent_id = ?", 0).Count(&total).Error
	if err != nil {
		global.Logger.Error("角色分页查询失败", zap.Error(err))
		return
	}
	if total < 0 {
		return
	}

	offset := (pagination.Page - 1) * pagination.Limit
	err = db.Where("parent_id = ?", 0).Offset(offset).Limit(pagination.Limit).Find(&roles).Error
	if err != nil {
		global.Logger.Error("角色分页查询失败", zap.Error(err))
		return
	}
	for k := range roles {
		err = s.findChildrenRole(&roles[k])
	}
	return
}

func (s *RoleService) findChildrenRole(role *system.SysRole) error {
	err := global.DB.Where("parent_id = ?", role.ID).Find(&role.Children).Error
	if len(role.Children) > 0 {
		for k := range role.Children {
			err = s.findChildrenRole(&role.Children[k])
		}
	}
	return err
}

func (s *RoleService) UpdateRole(r request.UpdateRole) error {
	role := system.SysRole{}
	role.ID = r.ID
	role.Name = r.Name
	role.Desc = r.Desc
	global.Logger.Error("abcccccc")
	err := global.DB.Where("id = ?", role.ID).Find(&system.SysRole{}).Updates(&role).Error
	return err
}

func (s *RoleService) DeleteRole(id uint) error {

	role := system.SysRole{}
	role.ID = id
	result := global.DB.Where("id = ?", id).Find(&role)
	if result.RowsAffected == 0 {
		return errors.New("不存在当前角色")
	}
	var total int64 = 0
	err := global.DB.Model(&system.SysRole{}).Where("parent_id = ?", role.ID).Count(&total).Error
	if err != nil {
		return errors.New("查询子角色错误")
	}
	if total > 0 {
		return errors.New("存在子角色，请删除子角色")
	}
	err = global.DB.Model(&system.SysRole{}).Where("id = ?", id).Delete(role).Error
	if err != nil {
		global.Logger.Error("删除角色失败", zap.Error(err))
		return errors.New("删除角色设备")
	}
	return nil
}
