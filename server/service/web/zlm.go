package web

import (
	"go.uber.org/zap"
	"nebula.xyz/global"
	"nebula.xyz/helper"
	"nebula.xyz/model/system"
	"time"
)

type ZLMService struct{}

func (z *ZLMService) UpdateServerStatus(server *system.MediaServer) {
	now := time.Now()
	server.KeepLiveAt = &now
	server.Status = helper.MediaStatusON
	if err := global.DB.Model(&system.MediaServer{}).Where("address = ?", server.Address).Updates(server).Error; err != nil {
		global.Logger.Error("媒体服务器状态更新失败", zap.Error(err))
		return
	}
}