package web

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"nebula.xyz/global"
	"nebula.xyz/helper"
	"nebula.xyz/model"
	"nebula.xyz/model/request"
	"nebula.xyz/model/system"
	"nebula.xyz/sip"
)

type VideoApi struct{}

// 点播
func (h *VideoApi) PlayVideo(ctx *gin.Context) {

	global.Logger.Info("执行PlayVideo函数")
	sip.Play(&system.Stream{
		ChannelId: "37070000081318000012",
		DeviceId:  "37070000081118000001",
	})
	model.OKWithMessage("Hello!!!", ctx)
}

// RecordVideo 开始录像
func (h *VideoApi) RecordVideo(ctx *gin.Context) {
	global.Logger.Info("执行RecordVideo函数")
	var record request.RecordVideo
	err := ctx.ShouldBindJSON(&record)
	if err != nil {
		global.Logger.Info("获取录像信息错误", zap.Error(err))
		model.ErrorWithMessage("获取信息错误", ctx)
		return
	}
	device := system.Device{DeviceId: record.DeviceId}
	if err := device.DeviceById(); err != nil {
		global.Logger.Info("设备ID错误")
		model.ErrorWithMessage("设备不存在", ctx)
		return
	}
	channel := system.DeviceChannel{ChannelId: record.ChannelId}
	if err := channel.DeviceChannelById(); err != nil {
		global.Logger.Info("设备通道不存在")
		model.ErrorWithMessage("设备通道不存在", ctx)
		return
	}
	if device.Status != helper.DeviceOnline || channel.Status != helper.ChannelStatusON {
		global.Logger.Info("设备或通道离线")
		model.ErrorWithMessage("设备或通道离线", nil)
		return
	}

	_ = videoService.RecordVideo(device, channel)
	if err != nil {
		return
	}
}
