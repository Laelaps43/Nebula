package web

import (
	"encoding/json"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"math"
	"nebula.xyz/global"
	"nebula.xyz/helper"
	"nebula.xyz/model/request"
	"nebula.xyz/model/request/zlm"
	"nebula.xyz/model/response"
	"nebula.xyz/model/system"
	"nebula.xyz/utils"
	"strings"
	"time"
)

type RecordServer struct{}

func (r *RecordServer) GetAllVideoRecord(page request.Pagination) (list []response.RecordPageResponse, total int64, err error) {
	db := global.DB.Debug()
	var recordTmpList []struct {
		Stream    string
		Duration  float64
		StartTime int64
	}
	err = db.Model(&system.Record{}).Group("stream").Count(&total).Error
	if err != nil {
		global.Logger.Error("查询总记录数失败", zap.Error(err))
		return nil, 0, err
	}
	offset := (page.Page - 1) * page.Limit
	err = db.Model(&system.Record{}).Select("stream, SUM(time_len) as duration, MAX(start_time) as start_time").
		Group("stream").
		Offset(offset).
		Limit(page.Limit).
		Scan(&recordTmpList).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, 0, nil
		}
		global.Logger.Error("查询视频记录失败", zap.Error(err))
		return nil, 0, err
	}
	var streamId []string
	for _, recordTmp := range recordTmpList {
		streamId = append(streamId, utils.HexToStream(recordTmp.Stream))
	}

	var streams []system.Stream
	err = db.Model(&system.Stream{}).Where("stream_id IN ?", streamId).Find(&streams).Error
	if err != nil {
		global.Logger.Error("查询视频流失败", zap.Error(err))
		return nil, 0, err
	}
	streamMap := make(map[string]system.Stream)
	channelIdList := make([]string, len(streamId))
	deviceIdList := make([]string, len(streamId))

	for _, stream := range streams {
		streamMap[stream.StreamId] = stream
		channelIdList = append(channelIdList, stream.ChannelId)
		deviceIdList = append(deviceIdList, stream.DeviceId)
	}
	var deviceList []system.Device
	err = db.Model(&system.Device{}).Where("device_id IN ?", deviceIdList).Find(&deviceList).Error
	if err != nil {
		global.Logger.Error("查询视频流失败", zap.Error(err))
		return nil, 0, err
	}

	deviceIdNameMap := make(map[string]string)
	for _, device := range deviceList {
		deviceIdNameMap[device.DeviceId] = device.Name
	}

	var channelList []system.DeviceChannel
	err = db.Model(&system.DeviceChannel{}).Where("channel_id IN ?", channelIdList).Find(&channelList).Error
	if err != nil {
		global.Logger.Error("查询视频流失败", zap.Error(err))
		return nil, 0, err
	}

	channelIdNameMap := make(map[string]string)

	for _, channel := range channelList {
		channelIdNameMap[channel.ChannelId] = channel.Name
	}

	resultList := make([]response.RecordPageResponse, 0)
	for _, recordTmp := range recordTmpList {
		stream := streamMap[utils.HexToStream(recordTmp.Stream)]
		tmp := response.RecordPageResponse{
			Stream:         utils.HexToStream(recordTmp.Stream),
			ChannelName:    channelIdNameMap[stream.ChannelId],
			ChannelID:      stream.ChannelId,
			DeviceName:     deviceIdNameMap[stream.DeviceId],
			IsRecording:    stream.Record,
			LastRecordTime: time.Unix(recordTmp.StartTime, 0).Format("2006-01-02 15:04:05"),
			Duration:       int64(recordTmp.Duration),
		}
		resultList = append(resultList, tmp)
	}
	return resultList, total, nil
}

func (r *RecordServer) GetVideoDateRange(rangeTime request.RecordRange) []string {
	timeFormat := "2006-01-02"
	startTime, err := time.Parse(timeFormat, rangeTime.Start)
	if err != nil {
		global.Logger.Debug("解析初始时间错误", zap.Error(err))
		return make([]string, 0)
	}
	endTime, err := time.Parse(timeFormat, rangeTime.End)
	if err != nil {
		global.Logger.Debug("解析结束时间错误", zap.Error(err))
		return make([]string, 0)
	}
	var recordDateList []string
	err = global.DB.Model(&system.Record{}).
		Select("record_date").
		Where("record_date >= ? AND record_date<= ? AND stream = ?", startTime, endTime, utils.StreamToHex(rangeTime.Stream)).
		Group("record_date").Scan(&recordDateList).Error
	if err != nil {
		global.Logger.Error("查询录像记录失败", zap.Error(err))
		return make([]string, 0)
	}
	return recordDateList
}

func (r *RecordServer) GetSelectRecord(stream string, selectTime time.Time) ([]response.RecordSelect, error) {
	var records []system.Record
	err := global.DB.Debug().Model(&system.Record{}).Where("stream = ? AND record_date = ? ", utils.StreamToHex(stream), selectTime.Format("2006-01-02")).Find(&records).Error
	if err != nil {
		global.Logger.Error("获取记录失败", zap.Error(err))
		return nil, err
	}
	timeFormat := "15:04:05"

	recordSelectList := make([]response.RecordSelect, 0)
	for _, record := range records {
		unix := time.Unix(record.StartTime, 0)
		beginTime := unix.Format(timeFormat)
		endUnix := unix.Add(time.Duration(math.Round(record.TimeLen)) * time.Second)
		endTime := endUnix.Format(timeFormat)
		recordSelectList = append(recordSelectList, response.RecordSelect{
			Label: beginTime + "-" + endTime,
			Value: record.ID,
		})
	}
	return recordSelectList, nil
}

func (r *RecordServer) GetRecordPlay(id string, stream string) (string, error) {

	var record system.Record
	err := global.DB.Model(&system.Record{}).Where("id = ? AND stream = ?", id, utils.StreamToHex(stream)).First(&record).Error
	if err != nil {
		global.Logger.Error("查询录像记录错误", zap.Error(err))
		return "", err
	}
	streamId := utils.GetRandStreamId()
	playRecord := &zlm.PlayRecord{
		Secret:     global.MediaServer.GetSecret(),
		Vhost:      record.Vhost,
		App:        record.App,
		Stream:     streamId,
		FilePath:   record.FilePath,
		FileRepeat: 0,
	}
	recordJson, _ := json.Marshal(playRecord)
	ZLMResponse, err := utils.ZLMHttpRequest(helper.ZlmLoadMP4File, strings.NewReader(string(recordJson)))
	if err != nil {
		global.Logger.Error("播放失败", zap.Error(err))
		return "", err
	}
	zlmRecord := response.ZLMRecord{}
	_ = json.Unmarshal(ZLMResponse, &zlmRecord)
	if zlmRecord.Code == 0 {
		return fmt.Sprintf("/rtp/%s/hls.m3u8", streamId), nil
	}
	return "", errors.New("播播放错误")
}
