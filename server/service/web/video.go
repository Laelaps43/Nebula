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

func (v *VideoService) StopPlay(stream *system.Stream) (err error) {
	// 判断流是否在录像，如果在录像，只将流的URL删掉
	record := &request.IsRecording{
		// TODO 需要更改
		Secret: global.MediaServer.GetSecret(),
		Type:   "1",
		Vhost:  "__defaultVhost__",
		App:    "rtp",
		Stream: utils.StreamToHex(stream.StreamId),
	}
	body, err := json.Marshal(record)
	if err != nil {
		global.Logger.Error("StopPlay 转换Json错误", zap.Error(err))
		return err
	}
	httpRequest, err := utils.ZLMHttpRequest(helper.ZlmIsRecording, strings.NewReader(string(body)))
	if err != nil {
		global.Logger.Info("ZLM isRecording 请求错误", zap.Error(err))
		return
	}
	global.Logger.Info("StopPlay获取ZLMediaKit回应", zap.String("response", httpRequest))
	return err
}
