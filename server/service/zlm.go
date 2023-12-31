package service

import (
	"go.uber.org/zap"
	"nebula.xyz/global"
	"nebula.xyz/model/request/zlm"
	"nebula.xyz/model/system"
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
