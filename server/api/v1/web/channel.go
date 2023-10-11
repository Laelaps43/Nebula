package web

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"nebula.xyz/global"
	"nebula.xyz/model"
	"nebula.xyz/model/request"
	resp "nebula.xyz/model/response"
	"nebula.xyz/model/system"
	"strconv"
)

type ChannelApi struct{}

// GetAllChannels 获取所有通道
func (c *ChannelApi) GetAllChannels(ctx *gin.Context) {
	var channel *system.DeviceChannel
	channels, err := channel.GetAllChannels()
	if err != nil {
		global.Logger.Error("查询通道失败", zap.Error(err))
		model.ErrorWithMessage("查询全部通道失败", ctx)
		return
	}
	model.OkWithDetailed(channels, "查询成功", ctx)
}

// GetChannelInfoById 根据通道Id获取通道信息
func (c *ChannelApi) GetChannelInfoById(ctx *gin.Context) {
	channelId := ctx.Param("channelId")
	if len(channelId) != 20 {
		model.ErrorWithMessage("通道Id错误", ctx)
		return
	}
	channel, err := channelService.GetChannelInfoById(channelId)
	if err != nil {
		model.ErrorWithMessage("获取通道信息错误", ctx)
		return
	}
	model.OkWithDetailed(channel, "获取通道信息成功", ctx)
}

// UpdateChannelInfo 根据通道ID更新通道信息
func (c *ChannelApi) UpdateChannelInfo(ctx *gin.Context) {
	var channel *request.ChannelReq
	err := ctx.ShouldBindJSON(&channel)
	if err != nil {
		model.ErrorWithMessage("绑定数据错误", ctx)
		return
	}
	if len(channel.ChannelId) != 20 || channel.Name == "" {
		model.ErrorWithMessage("参数错误", ctx)
		return
	}
	deviceChannel := &system.DeviceChannel{
		ChannelId: channel.ChannelId,
		Name:      channel.Name,
	}
	err = channelService.UpdateChannelInfoById(deviceChannel)
	if err != nil {
		model.ErrorWithMessage("更新失败", ctx)
		return
	}
	model.OKWithMessage("更新成功", ctx)
}

// GenerateChannel 生成通道
func (c *ChannelApi) GenerateChannel(ctx *gin.Context) {
	var gen *request.ChannelGenerate

	err := ctx.ShouldBindJSON(&gen)
	if err != nil {
		model.ErrorWithMessage("绑定数据错误", ctx)
		return
	}
	if len(gen.DeviceId) != 20 {
		model.ErrorWithMessage("数据错误", ctx)
		return
	}
	num, _ := strconv.Atoi(gen.ChannelNum)
	if num <= 0 {
		return
	}
	device, err := deviceService.GetDeviceInfoById(gen.DeviceId)
	if err != nil {
		global.Logger.Error("设备不存在", zap.Error(err))
		model.ErrorWithMessage("设备不存在", ctx)
		return
	}
	channels, err := channelService.GenerateChannel(num, device)

	channel := resp.GenerateInfo{
		Device:     device,
		Channels:   channels,
		ChannelSum: len(channels),
	}
	model.OkWithDetailed(channel, "生成通道成功", ctx)
}
