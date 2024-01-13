package system

type SysRoleButton struct {
	SysRoleId   uint `gorm:"column:sys_role_id"`
	SysButtonId uint `gorm:"column:sys_button_id"`
}

func (s *SysRoleButton) TableName() string {
	return "sys_role_buttons"
}
