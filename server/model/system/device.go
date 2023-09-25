package system

import (
	"time"

	"nebula.xyz/model"
)

type Device struct {
	model.NEBULA
	DeviceId     string    `gorm:"comment:设备国标ID"` // 设备国标Id
	Name         string    `gorm:"comment:设备名称"`   // 设备名称
	Manufacture  string    `gorm:"comment:制造商"`    // 制造商
	Transport    string    `gorm:"comment:传输协议"`   // 传输协议
	StreamModel  string    `gorm:"comment:流传输模式"`  // 流传输模式
	IP           string    `gorm:"comment:设备地址"`   // 设备地址
	Port         int       `gorm:"comment:设备端口"`   // 设备端口
	RegisterTime time.Time `gorm:"comment:设备注册时间"` // 设置注册时间
	KeepLiveTime int64     `gorm:"comment:心跳时间"`   // 心跳时间
	ChannelCount int       `gorm:"comment:通道个数"`   // 通道个数
}

func (d Device) TableName() string {
	return "device"
}
