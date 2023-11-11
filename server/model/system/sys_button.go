package system

import "nebula.xyz/model"

// SysButton Button 系统按钮
type SysButton struct {
	model.NEBULA
	ID   uint   `gorm:"primaryKey comment:主键"`
	Slug string `gorm:"comment:按钮唯一表示"`
	Name string `gorm:"comment:按钮名称"`
	Desc string `gorm:"按钮描述"`
}

func (b *SysButton) TableName() string {
	return "sys_button"
}
