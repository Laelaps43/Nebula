package sip

// Sip 服务

import (
	"errors"
	"fmt"
	"github.com/ghettovoice/gosip"
	l "github.com/ghettovoice/gosip/log"
	"github.com/ghettovoice/gosip/sip"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"nebula.xyz/global"
	"nebula.xyz/helper"
	"nebula.xyz/model/system"
	"sync"
)

// Server Sip Server
var (
	Server    gosip.Server
	SyncOnce  sync.Once
	sipServer *system.SipServer
)

func NewSipServer() {
	SyncOnce.Do(func() {
		Server = gosip.NewServer(
			gosip.ServerConfig{
				Host:      sipServer.IP,
				UserAgent: sipServer.UserAgent,
			},
			nil,
			nil,
			// TODO 将Logger 换成 zap
			l.NewDefaultLogrusLogger(),
		)
	})

	if err := Server.Listen("tcp", fmt.Sprintf("%s:%d", sipServer.IP, sipServer.Port)); err != nil {
		global.Logger.Error("SIP服务器监听tcp协议失败！", zap.Error(err))
		return
	}
	global.Logger.Info("SIP服务器正在监听tcp协议...")
	if err := Server.Listen("udp", fmt.Sprintf("%s:%d", sipServer.IP, sipServer.Port)); err != nil {
		global.Logger.Error("SIP服务器监听udp协议失败！", zap.Error(err))
		return
	}
	global.Logger.Info("SIP服务器正在监听tcp协议...")
	registerHandler()
}

func InitSipServer() {
	result := global.DB.Order("sort asc").Where("status = ?", helper.SipServerON).First(&sipServer)
	// 数据库中可以找到一条数据，则使用数据库中查找到的，负责使用配置文件
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		global.Logger.Debug("未从数据库中获取到服务器信息，从配置文件中获取。")
		sipServer.SipId = global.CONFIG.SIP.SipId
		sipServer.Port = global.CONFIG.SIP.Port
		sipServer.IP = global.CONFIG.SIP.Ip
		sipServer.Realm = global.CONFIG.SIP.Realm
		sipServer.Password = global.CONFIG.SIP.Password
		sipServer.UserAgent = global.CONFIG.SIP.UserAgent
		sipServer.DevicePrefix = global.CONFIG.SIP.DevicePrefix
		sipServer.ChannelPrefix = global.CONFIG.SIP.ChannelPrefix
		sipServer.Status = helper.SipServerON
		sipServer.Sort = 1
		global.DB.Create(&sipServer)
	}
	global.Logger.Info("初始化SIP服务中...")
	NewSipServer()
}

func registerHandler() {
	_ = Server.OnRequest(sip.REGISTER, Register)
	_ = Server.OnRequest(sip.MESSAGE, Message)
}
