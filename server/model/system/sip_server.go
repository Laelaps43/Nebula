package system

// sip实体

import (
	"nebula.xyz/global"
	"nebula.xyz/helper"
	"nebula.xyz/model"
)

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

func (s *SipServer) TableName() string {
	return "sip_server"
}

// GetSipServerById 根据sip_id 获取sipServer
func (s *SipServer) GetSipServerById() (err error) {
	if err = global.DB.Model(&SipServer{}).Where("sip_id = ?", s.SipId).First(&s).Error; err != nil {
		return nil
	}
	return err
}

// GetSipServerOnLine 获取设备在线Sip服务器
func (s *SipServer) GetSipServerOnLine() (err error) {
	if err = global.DB.Model(&SipServer{}).Order("sort asc").Where("status = ?", helper.SipServerON).First(&s).Error; err != nil {
		return err
	}
	return nil
}
