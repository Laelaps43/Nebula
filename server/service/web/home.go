package web

import (
	"go.uber.org/zap"
	"nebula.xyz/global"
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
	err := global.DB.Model(&system.Stream{}).Where("record = ?", helper.StreamRecorded).Count(&videoCount).Error
	if err != nil {
		global.Logger.Error("获取录像设备数量错误", zap.Error(err))
		videoCount = 0
	}
	result := resp.OverViewResult{
		OnlineDevice:  onLineCount,
		OfflineDevice: offLineCount,
		Channel:       channelCount,
		Video:         videoCount,
	}
	return result
}
