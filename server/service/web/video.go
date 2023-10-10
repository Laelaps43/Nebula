package web

import (
	"encoding/json"
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"nebula.xyz/global"
	"nebula.xyz/helper"
	"nebula.xyz/model/request"
	"nebula.xyz/model/system"
	"nebula.xyz/sip"
	"nebula.xyz/utils"
	"strings"
)

type VideoService struct{}

func (v *VideoService) RecordVideo(device system.Device, channel system.DeviceChannel) error {
	stream := &system.Stream{DeviceId: device.DeviceId, ChannelId: channel.ChannelId}
	err := stream.GetStreamByDeviceAndChannel()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		stream, err = sip.Play(stream)
		if err != nil {
			global.Logger.Error("录像错误", zap.Error(err))
			return err
		}
	}
	// 向ZLM提供请求录像
	start := &request.StartRecord{
		Secret:         global.MediaServer.GetSecret(),
		Type:           "1",
		Vhost:          "__defaultVhost__",
		App:            "rtp",
		Stream:         utils.StreamToHex(stream.StreamId),
		MaxSecond:      60,
		CustomizedPath: "/opt/media/video",
	}
	body, err := json.Marshal(start)
	if err != nil {
		global.Logger.Error("转换StartRecord错误")
	}
	global.Logger.Error(string(body))
	httpRequest, err := utils.ZLMHttpRequest(helper.ZlmStartRecord, strings.NewReader(string(body)))
	global.Logger.Error(httpRequest)
	return nil
}
