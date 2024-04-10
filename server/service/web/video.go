package web

import (
	"encoding/json"
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"nebula.xyz/global"
	"nebula.xyz/helper"
	"nebula.xyz/model/request"
	"nebula.xyz/model/request/zlm"
	"nebula.xyz/model/response"
	"nebula.xyz/model/system"
	"nebula.xyz/utils"
	"strings"
)

type VideoService struct{}

func (v *VideoService) RecordVideo(device system.Device, channel system.DeviceChannel) error {
	stream := &system.Stream{DeviceId: device.DeviceId, ChannelId: channel.ChannelId}
	err := stream.GetStreamByDeviceAndChannel()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			global.Logger.Error("流未找到, 尝试拉流")
		} else {
			global.Logger.Error("记录失败", zap.Error(err))
			return errors.New("播放错误")
		}
	}
	if stream.Status == helper.StreamClose {
		global.Logger.Debug("流不在线，请先获取流")
		return errors.New("流不在线")
	}

	//// 判断流是否在录制
	//isRecord := &zlm.IsRecording{
	//	Secret: global.MediaServer.GetSecret(),
	//	Type:   helper.RecordMP4,
	//	Vhost:  stream.VHost,
	//	App:    stream.App,
	//	Stream: utils.StreamToHex(stream.StreamId),
	//}
	//marshal, _ := json.Marshal(isRecord)
	//httpResponse, err := utils.ZLMHttpRequest(helper.ZlmIsRecording, strings.NewReader(string(marshal)))
	//recordResponse := response.ZLMRecord{}
	//err = json.Unmarshal(httpResponse, &recordResponse)
	//if err != nil {
	//	global.Logger.Error("Json解析录制返回值错误", zap.Error(err))
	//	return errors.New("录制错误")
	//}
	if stream.Record == helper.StreamRecorded {
		global.Logger.Error("流已被录制")
		return errors.New("流已被录制")
	}

	// 开始录制
	start := &zlm.StartRecord{
		Secret:         global.MediaServer.GetSecret(),
		Type:           helper.RecordMP4,
		Vhost:          stream.VHost,
		App:            stream.App,
		Stream:         utils.StreamToHex(stream.StreamId),
		MaxSecond:      helper.RecordMaxSecond,
		CustomizedPath: global.CONFIG.Media.RecordPath,
	}
	body, err := json.Marshal(start)
	if err != nil {
		global.Logger.Error("转换StartRecord错误")
		return errors.New("录制错误")
	}
	global.Logger.Error(string(body))
	httpResponse, err := utils.ZLMHttpRequest(helper.ZlmStartRecord, strings.NewReader(string(body)))
	if err != nil {
		global.Logger.Error("发送录制请求错误", zap.Error(err))
		return errors.New("录制错误")
	}
	recordResponse := response.ZLMRecord{}
	err = json.Unmarshal(httpResponse, &recordResponse)
	if err != nil {
		global.Logger.Error("Json解析录制返回值错误", zap.Error(err))
		return errors.New("录制错误")
	}
	if recordResponse.Result {
		stream.Record = helper.StreamRecorded
		global.DB.Where("stream_id = ?", stream.StreamId).Save(stream)
		return nil
	}
	return errors.New("录制错误")
}

func (v *VideoService) StopPlay(stream *system.Stream) (err error) {
	// 判断流是否在录像，如果在录像，只将流的URL删掉
	record := &zlm.IsRecording{
		// TODO 需要更改
		Secret: global.MediaServer.GetSecret(),
		Type:   1,
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
	global.Logger.Info("StopPlay获取ZLMediaKit回应", zap.String("response", string(httpRequest)))
	return err
}

func (v *VideoService) StopRecord(record request.RecordVideo) error {

	stream := system.Stream{}

	err := global.DB.Model(&stream).Where("channel_id = ? And device_id = ?", record.ChannelId, record.DeviceId).First(stream).Error
	if err != nil {
		global.Logger.Debug("获取流数据失败", zap.Error(err))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("流不存在")
		}
		return errors.New("停止录制错误")
	}
	if stream.Record == helper.StreamUnRecord {
		return errors.New("当前设备未在录制")
	}

	stopRecord := zlm.StopRecord{
		Secret: global.MediaServer.GetSecret(),
		Type:   helper.RecordMP4,
		Vhost:  stream.VHost,
		App:    stream.App,
		Stream: utils.StreamToHex(stream.StreamId),
	}
	marshal, _ := json.Marshal(stopRecord)
	ZLMResponse, err := utils.ZLMHttpRequest(helper.ZlmStopRecording, strings.NewReader(string(marshal)))
	if err != nil {
		global.Logger.Error("发送停止录制设备", zap.Error(err))
		return errors.New("停止录制失败")
	}
	zlmRecord := response.ZLMRecord{}
	err = json.Unmarshal(ZLMResponse, &zlmRecord)
	if err != nil {
		global.Logger.Error("Json解析录制返回值错误", zap.Error(err))
		return errors.New("停止录制错误")
	}
	if zlmRecord.Result {
		stream.Record = helper.StreamUnRecord
		global.DB.Save(stream)
		return nil
	}
	return errors.New("停止录制错误")
}
