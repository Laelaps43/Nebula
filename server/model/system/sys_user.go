package system

import "nebula.xyz/model"

type SysUser struct {
	model.NEBULA
	ID        uint      `gorm:"primaryKey"` // 主键ID
	UserName  string    `json:"username" gorm:"comment:用户登录名"`
	PassWord  string    `json:"-" gorm:"comment:用户登录密码"`
	HeaderImg string    `json:"headerImg" gorm:"comment:用户头像"`
	Email     string    `json:"email" gorm:"comment:用户邮箱"`
	Enable    int       `json:"enable" gorm:"comment:用户是否被冻结 1 正常 0 冻结"`
	RoleID    uint      `gorm:"comment:用户角色ID"`
	Role      SysRole   `gorm:"foreignKey:RoleID"`
	Roles     []SysRole `gorm:"many2many:sys_user_roles;"` // 用户可以有多个角色
}

func (u SysUser) TableName() string {
	return "sys_user"
}
