package system

import "nebula.xyz/model"

// SysMenu Menu 系统菜单
type SysMenu struct {
	model.NEBULA
	ID       uint   `gorm:"primaryKey comment:主键"`
	Name     string `gorm:"comment:菜单名称"`
	ParentId string `gorm:"comment:父菜单ID"`
	Slug     string `gorm:"comment:菜单唯一表示"`
	Desc     string `gorm:"comment:菜单描述"`
}

func (m *SysMenu) TableName() string {
	return "sys_menu"
}
