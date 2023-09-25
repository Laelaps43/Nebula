package system

import "nebula.xyz/model"

type DeviceChannel struct {
	model.NEBULA
	ChannelId string `gorm:"comment:设备通道编号"` // 通道目标编号
	DeviceId  int    `gorm:"comment:设备ID"`   // 设备Id
	Device    Device
}

func (DeviceChannel) TableName() string {
	return "device_channel"
}
