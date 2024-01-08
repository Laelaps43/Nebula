package service

import (
	"go.uber.org/zap"
	"nebula.xyz/global"
	"nebula.xyz/helper"
	"nebula.xyz/model/request/zlm"
	"nebula.xyz/model/system"
	"nebula.xyz/utils"
	"time"
)

type ZLMService struct{}

// UpdateServerStatus 更新服务器状态
func (z *ZLMService) UpdateServerStatus(keepalive zlm.ServerKeepalive) {
	now := time.Now()
	server := &system.MediaServer{
		MediaServerId: keepalive.MediaServerId,
		KeepLiveAt:    &now,
		MediaSource:   keepalive.Data.MediaSource,
		TcpServer:     keepalive.Data.TcpServer,
		TcpClient:     keepalive.Data.TcpClient,
		TcpSession:    keepalive.Data.TcpSession,
		UdpServer:     keepalive.Data.UdpServer,
		UdpSession:    keepalive.Data.UdpSession,
	}
	if err := global.DB.Model(&system.MediaServer{}).Where("media_server_id = ?", server.MediaServerId).Updates(server).Error; err != nil {
		global.Logger.Error("媒体服务器状态更新失败", zap.Error(err))
		return
	}
}

// UpdateStreamInfo 更新流信息
func (z *ZLMService) UpdateStreamInfo(steam *system.Stream) error {
	return steam.Update()
}

func (z *ZLMService) OnStreamChanged(change zlm.StreamChange) error {
	stream := system.Stream{}
	err := global.DB.Model(&system.Stream{}).Where("stream_id = ?", utils.HexToStream(change.Stream)).Find(&stream).Error
	if err != nil {
		global.Logger.Error("流注册或注销失败", zap.Error(err))
		return err
	}
	if change.Regist {
		global.Logger.Info("收到流注册", zap.String("流Id", change.Stream))
		// 注册
		stream.Status = helper.StreamStart
		stream.OriginType = change.OriginType
		stream.TotalReaderCount = change.TotalReaderCount
		stream.Schema = change.Schema
		stream.VHost = change.Vhost
		stream.StreamType = change.OriginTypeStr
	} else {
		// 注销
		global.Logger.Info("收到流注销", zap.String("流Id", change.Stream))
		stream.Status = helper.StreamClose
		global.Logger.Debug("stream值", zap.Any("stream", stream))
	}
	err = global.DB.Model(&system.Stream{}).Where("stream_id = ? ", utils.HexToStream(change.Stream)).Updates(stream).Error
	if err != nil {
		global.Logger.Error("更新流状态错误", zap.Error(err))
		return err
	}
	return nil
}

func (z *ZLMService) OnRecordMp4(mp4 zlm.RecordMp4) error {
	timeObj := time.Unix(mp4.StartTime, 0)
	record := system.Record{
		MediaServerId: mp4.MediaServerId,
		App:           mp4.App,
		FileName:      mp4.FileName,
		FilePath:      mp4.FilePath,
		FileSize:      mp4.FileSize,
		Folder:        mp4.Folder,
		StartTime:     mp4.StartTime,
		Stream:        mp4.Stream,
		TimeLen:       mp4.TimeLen,
		URL:           mp4.URL,
		Vhost:         mp4.Vhost,
		RecordDate:    timeObj.Format("2006-01-02"),
	}
	err := global.DB.Model(&system.Record{}).Save(&record).Error
	if err != nil {
		global.Logger.Error("记录MP4失败", zap.Error(err))
		return err
	}
	return nil
}
