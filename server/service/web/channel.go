package web

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"nebula.xyz/global"
	"nebula.xyz/helper"
	"nebula.xyz/model/system"
	"nebula.xyz/utils"
)

type ChannelService struct{}

// GenerateChannel 生成指定通道个数
func (c *ChannelService) GenerateChannel(n int, device *system.Device) (channels []*system.DeviceChannel, err error) {
	channels = make([]*system.DeviceChannel, 0, n)
	server := system.SipServer{}
	err = server.GetSipServerOnLine()
	if err != nil {
		global.Logger.Error("获取SipServer信息错误", zap.Error(err))
		return nil, errors.New("获取SipServer信息错误")
	}
	// 生成通道ID
	// 生成通道规则为，前3为设备ID的最后3为，后3为随机生成
	prex := server.ChannelPrefix
	for n > 0 {
		b := device.DeviceId[17:20]
		a := utils.RandInt(3)
		key := fmt.Sprintf("%s%s%s", prex, b, a)
		cha := system.DeviceChannel{
			ChannelId: key,
		}
		if !cha.IsExist() {
			cha.Status = helper.ChannelStatusOFF
			cha.DeviceId = device.DeviceId
			// 名称为最后6位
			cha.Name = fmt.Sprintf("%s%s", b, a)
			err := cha.ChannelAdd()
			if err != nil {
				global.Logger.Error("添加通道错误")
			} else {
				n--
				channels = append(channels, &cha)
			}
		}
	}
	return
}
