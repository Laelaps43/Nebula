package web

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"nebula.xyz/global"
	"nebula.xyz/model"
	"nebula.xyz/model/request"
	resp "nebula.xyz/model/response"
	"nebula.xyz/model/system"
)

type ChannelApi struct{}

// GetAllChannels 获取所有通道
func (c *ChannelApi) GetAllChannels(ctx *gin.Context) {
	var pagination request.Pagination

	err := ctx.ShouldBindJSON(&pagination)
	if err != nil {
		model.ErrorWithMessage("请检查分页数据", ctx)
		return
	}
	channelPagination, total, err := channelService.GetChannelPagination(pagination)
	if err != nil {
		global.Logger.Error("查询设备信息失败")
		model.ErrorWithMessage("获取失败", ctx)
		return
	}
	model.OkWithDetailed(resp.PaginationResult{
		List:  channelPagination,
		Total: total,
	}, "获取成功", ctx)
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
	channel, err := channelService.GenerateChannel()
	if err != nil {
		model.ErrorWithMessage("生成设备失败", ctx)
		return
	}
	model.OkWithDetailed(channel.ChannelId, "生成设备成功", ctx)
	return
}

func (c *ChannelApi) CreateChannel(ctx *gin.Context) {

	var createChannel request.CreateChannel

	err := ctx.ShouldBindJSON(&createChannel)
	if err != nil {
		model.ErrorWithMessage("服务器内部异常", ctx)
		ctx.Abort()
	}

	err = channelService.CreateChannel(createChannel)
	if err != nil {
		global.Logger.Error("创建通道失败", zap.Error(err))
		model.ErrorWithMessage("创建通道失败", ctx)
		ctx.Abort()
	}
	model.OKWithMessage("创建通道成功", ctx)
}

func (c *ChannelApi) UpdateChannel(ctx *gin.Context) {

	var channel system.DeviceChannel
	err := ctx.ShouldBindJSON(&channel)
	if err != nil {
		model.ErrorWithMessage(err.Error(), ctx)
		return
	}
	if len(channel.ChannelId) != 20 || len(channel.Name) == 0 {
		model.ErrorWithMessage("参数错误", ctx)
	}

	err = channelService.UpdateChannelInfoById(&channel)
	if err != nil {
		model.ErrorWithMessage("更新错误", ctx)
		return
	}
	model.OKWithMessage("更新成功", ctx)
}

func (c *ChannelApi) DeleteChannel(ctx *gin.Context) {
	channelId := ctx.Param("channelId")

	if len(channelId) != 20 {
		model.ErrorWithMessage("通道Id错误", ctx)
		return
	}

	err := channelService.DeleteChannel(channelId)
	if err != nil {
		model.ErrorWithMessage("删除通道失败", ctx)
		return
	}
	model.OK(ctx)
}
