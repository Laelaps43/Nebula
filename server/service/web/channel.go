package web

import (
	"errors"
	"go.uber.org/zap"
	"nebula.xyz/global"
	"nebula.xyz/helper"
	"nebula.xyz/model/request"
	"nebula.xyz/model/system"
	"nebula.xyz/utils"
)

type ChannelService struct{}

func (c *ChannelService) GetChannelPagination(pagination request.Pagination) (channels []system.DeviceChannel, total int64, err error) {
	db := global.DB.Model(&system.DeviceChannel{})
	err = db.Where("device_id = ?", pagination.DeviceId).Count(&total).Error
	if err != nil {
		global.Logger.Error("设备通道分页查询失败", zap.Error(err))
		return
	}
	if total < 0 {
		return
	}
	offset := (pagination.Page - 1) * pagination.Limit
	err = db.Where("device_id = ?", pagination.DeviceId).Offset(offset).Limit(pagination.Limit).Find(&channels).Error
	if err != nil {
		global.Logger.Error("设备通道分页查询失败", zap.Error(err))
	}
	return
}

// GenerateChannel 生成指定通道个数
func (c *ChannelService) GenerateChannel() (channel *system.DeviceChannel, err error) {
	channel = &system.DeviceChannel{}
	server := system.SipServer{}
	err = server.GetSipServerOnLine()
	if err != nil {
		global.Logger.Error("获取SipServer信息错误", zap.Error(err))
		return nil, errors.New("获取SipServer信息错误")
	}
	// 生成设备ID
	for {
		randInt := utils.RandInt(6)
		key := server.ChannelPrefix + randInt
		channel.ChannelId = key
		exist := channel.IsExist()
		if !exist {
			break
		}
	}
	return channel, nil
}

// GetChannelInfoById 根据通道Id获取通道信息
func (c *ChannelService) GetChannelInfoById(id string) (channel *system.DeviceChannel, err error) {
	channel = &system.DeviceChannel{
		ChannelId: id,
	}
	err = channel.DeviceChannelById()
	if err != nil {
		return nil, err
	}
	return
}

func (c *ChannelService) UpdateChannelInfoById(channel *system.DeviceChannel) error {
	err := channel.UpdateChannelInfoById()
	return err
}

func (c *ChannelService) CreateChannel(create request.CreateChannel) error {
	channel := system.DeviceChannel{}
	tx := global.DB.Begin()
	result := tx.Where("channel_id = ?", create.ChannelId).Find(&channel)
	if result.RowsAffected != 0 {
		global.Logger.Info("通道id已存在", zap.String("channel_id", create.ChannelId))
		tx.Rollback()
		return result.Error
	}
	channel.ChannelId = create.ChannelId
	channel.Name = create.Name
	channel.DeviceId = create.DeviceId
	channel.Transport = helper.DefaultTransPort
	channel.Status = helper.ChannelStatusOFF
	createResult := tx.Create(&channel)
	if createResult.Error != nil {
		global.Logger.Error("创建通道失败", zap.Error(createResult.Error))
		tx.Rollback()
		return createResult.Error
	}
	tx.Commit()
	return nil
}

func (c *ChannelService) DeleteChannel(channelId string) error {

	channel := system.DeviceChannel{}
	channel.ChannelId = channelId

	err := channel.DeleteChannelById()
	if err != nil {
		global.Logger.Error("删除通道失败")
		return err
	}
	return nil
}
