package system

import "nebula.xyz/model"

// SysRole Role 系统角色
type SysRole struct {
	model.NEBULA
	ID        uint        `gorm:"not null;unique;primary_key;comment:角色ID;size:90"`
	Name      string      `gorm:"comment:角色名"`
	ParentId  *uint       `gorm:"comment:父角色ID"`
	Slug      string      `gorm:"comment:唯一标识"`
	Desc      string      `gorm:"comment:角色描述"`
	SysMenus  []SysMenu   `gorm:"many2many:sys_role_menus"`   // 角色可以有多个菜单
	SysButton []SysButton `gorm:"many2many:sys_role_buttons"` // 角色可以多个按钮
}

func (r *SysRole) TableName() string {
	return "sys_role"
}
