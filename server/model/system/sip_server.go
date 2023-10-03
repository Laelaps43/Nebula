package system

// sip实体

import "nebula.xyz/model"

type SipServer struct {
	model.NEBULA

	SipId string `gorm:"PrimaryKey, comment:SIP服务器ID"`
	IP    string `gorm:"comment:SIP服务器地址"`
	Port  uint   `gorm:"comment:SIP服务器端口"`
	Realm string `gorm:"comment:SIP区域"`
	// TODO 是否需要加密
	Password      string `gorm:"comment:SIP密码"`
	UserAgent     string `gorm:"comment:SIP用户代理"`
	DevicePrefix  string `gorm:"comment:设备前缀"`
	ChannelPrefix string `gorm:"comment:通道前缀"`
	Status        uint   `gorm:"comment:是否开启 1 开启，0 关闭"`
	Sort          uint   `gorm:"comment:排序"`
}
