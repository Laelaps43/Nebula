package system

import "nebula.xyz/model"

type User struct {
	model.NEBULA
	ID        int    `gorm:"primaryKey"` // 主键ID
	UserName  string `json:"username" gorm:"comment:用户登录名"`
	PassWord  string `json:"-" gorm:"comment:用户登录密码"`
	HeaderImg string `json:"headerImg" gorm:"comment:用户头像"`
	Email     string `json:"email" gorm:"comment:用户邮箱"`
	Enable    int    `json:"enable" gorm:"comment:用户是否被冻结 1 正常 0 冻结"`
}

func (u User) TableName() string {
	return "user"
}
