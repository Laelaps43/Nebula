package system

import "nebula.xyz/model"


type DeviceChannelId struct{
	model.NEBULA
	ChannelId			string				// 通道目标编号
	DeviceId			Device				// 设备Id
}