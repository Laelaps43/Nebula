package web

import (
	"nebula.xyz/helper"
	resp "nebula.xyz/model/response"
	"nebula.xyz/model/system"
)

type HomeService struct{}

// GetOverView 获取首页数据总览
func (s HomeService) GetOverView() resp.OverViewResult {

	device := system.Device{}
	// 获取在线设备
	device.Status = helper.DeviceOnline
	onLineCount := device.DeviceCountByStatus()
	// 获取离线设备
	device.Status = helper.DeviceOffline
	offLineCount := device.DeviceCountByStatus()

	// 获取通道
	channel := system.DeviceChannel{}
	channelCount := channel.ChannelCount()

	// 获取正在录像设备
	var videoCount int64 = 0
	result := resp.OverViewResult{
		OnlineDevice:  onLineCount,
		OfflineDevice: offLineCount,
		Channel:       channelCount,
		Video:         videoCount,
	}
	return result
}
