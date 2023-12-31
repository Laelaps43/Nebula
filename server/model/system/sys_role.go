package system

import (
	"nebula.xyz/global"
	"nebula.xyz/model"
)

// SysRole Role 系统角色
type SysRole struct {
	model.NEBULA
	ID        uint        `gorm:"not null;unique;primary_key;comment:角色ID;size:90" json:"id"`
	Name      string      `gorm:"comment:角色名" json:"name"`
	ParentId  *uint       `gorm:"comment:父角色ID" json:"-"`
	Slug      string      `gorm:"comment:唯一标识" json:"slug"`
	Desc      string      `gorm:"comment:角色描述" json:"desc"`
	SysMenus  []SysMenu   `gorm:"many2many:sys_role_menus" json:"-"`   // 角色可以有多个菜单
	SysButton []SysButton `gorm:"many2many:sys_role_buttons" json:"-"` // 角色可以多个按钮
	Children  []SysRole   `gorm:"-" json:"children,omitempty"`
}

func (r *SysRole) TableName() string {
	return "sys_role"
}

func (r *SysRole) GetRoleBySlug() error {
	result := global.DB.Where("slug = ?", r.Slug).First(&r)
	if result.RowsAffected != 0 {
		return result.Error
	}
	return nil
}

func (r *SysRole) SaveRole() error {
	err := global.DB.Save(r).Error
	return err
}
