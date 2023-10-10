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
	// 信令服务器设置
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
	// 初始化媒体服务器
	meidaTmp := &system.MediaServer{}
	result = global.DB.Order("sort asc").Where("status = ?", helper.MediaStatusON).First(&meidaTmp)
	global.MediaServer = meidaTmp
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		global.Logger.Info("未从数据库中获取到媒体服务器信息，从配置文件中获取。")
		global.MediaServer.SetAddress(global.CONFIG.Media.Address)
		global.MediaServer.SetRTSPPort(global.CONFIG.Media.RTSPPort)
		global.MediaServer.SetRestful(global.CONFIG.Media.Restful)
		global.MediaServer.SetRTP(global.CONFIG.Media.RTP)
		global.MediaServer.SetRTMPPort(global.CONFIG.Media.RTMPPort)
		global.MediaServer.SetSecret(global.CONFIG.Media.Secret)
		global.DB.Create(&meidaTmp)
	}
	//_, err := utils.ZLMHttpRequest(helper.ZlmGetApiList, nil)
	//if err != nil {
	//	global.Logger.Error("初始化媒体服务器失败")
	//return
	//}
	global.Logger.Info("初始化媒体服务器完成")
	NewSipServer()
}

func registerHandler() {
	_ = Server.OnRequest(sip.REGISTER, Register)
	_ = Server.OnRequest(sip.MESSAGE, Message)
}
