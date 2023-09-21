package system

import (
	"time"

	"nebula.xyz/model"
)

type Device struct{
	model.NEBULA
	DeviceId		string				// 设备国标Id
	Name			string				// 设备名称
	Manufacture		string				// 制造商
	Transposrt		string				// 传输协议
	StreamModel		string				// 流传输模式
	IP				string				// 设备地址
	Port			int					// 设备端口
	RegisterTime	time.Time			// 设置注册时间
	KeepLiveTime	int64				// 心跳时间
	ChannelCount	int					// 通道个数
}

func (d Device) TableName() string{
	return "device"
}