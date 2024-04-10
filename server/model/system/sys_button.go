package system

import "nebula.xyz/model"

// SysButton Button 系统按钮
type SysButton struct {
	model.NEBULA
	ID     uint   `gorm:"primaryKey comment:主键"`
	Slug   string `gorm:"uniqueIndex,comment:按钮唯一表示"`
	Name   string `gorm:"comment:按钮名称"`
	Desc   string `gorm:"comment:按钮描述"`
	Path   string `gorm:"comment:按钮对应URL"`
	Method string `gorm:"comment:请求方法"`
}

func (b *SysButton) TableName() string {
	return "sys_button"
}
