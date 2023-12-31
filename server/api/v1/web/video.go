package web

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"nebula.xyz/global"
	"nebula.xyz/helper"
	"nebula.xyz/model"
	"nebula.xyz/model/request"
	"nebula.xyz/model/response"
	"nebula.xyz/model/system"
	"nebula.xyz/sip"
)

type VideoApi struct{}

// PlayVideo 点播
func (v *VideoApi) PlayVideo(ctx *gin.Context) {
	var videoPayload request.VideoRequestPayload
	err := ctx.ShouldBindJSON(&videoPayload)
	if err != nil {
		model.ErrorWithMessage("参数错误, 播放错误", ctx)
		return
	}
	global.Logger.Info("执行PlayVideo函数")
	playStream, err := sip.Play(videoPayload)
	if err != nil {
		global.Logger.Error("点播错误:", zap.Error(err))
		model.ErrorWithMessage("播放错误", ctx)
		return
	}

	model.OkWithDetailed(
		response.PlayResponsePayload{
			HTTP:  playStream.HTTP,
			RTSP:  playStream.RTSP,
			RTMP:  playStream.RTMP,
			WSFLV: playStream.WSFLV,
		}, "点播成功", ctx)
	return
}

// StopPlay 停止直播，这里需要考虑流是否在录像
func (v *VideoApi) StopPlay(ctx *gin.Context) {
	var stop *request.StopPlay
	err := ctx.ShouldBindJSON(&stop)
	if err != nil {
		model.ErrorWithMessage("StopPlay绑定时间失败", ctx)
		return
	}
	stream := &system.Stream{
		DeviceId:  stop.DeviceId,
		ChannelId: stop.ChannelId,
	}
	err = stream.GetStreamByDeviceAndChannel()
	if err != nil {
		model.ErrorWithMessage("没有找到对应的流", ctx)
		return
	}
	videoService.StopPlay(stream)

}

// RecordVideo 开始录像
func (v *VideoApi) RecordVideo(ctx *gin.Context) {
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
