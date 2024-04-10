package system

type SysRoleMenus struct {
	SysRoleId uint `gorm:"column:sys_role_id"`
	SysMenuId uint `gorm:"column:sys_menu_id"`
}

func (s *SysRoleMenus) TableName() string {
	return "sys_role_menus"
}
