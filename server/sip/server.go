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
	"nebula.xyz/utils"
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
	// 信令服务器设置
	sipServer = &system.SipServer{
		SipId:         global.CONFIG.SIP.SipId,
		Port:          global.CONFIG.SIP.Port,
		IP:            global.CONFIG.SIP.Ip,
		Realm:         global.CONFIG.SIP.Realm,
		Password:      global.CONFIG.SIP.Password,
		UserAgent:     global.CONFIG.SIP.UserAgent,
		DevicePrefix:  global.CONFIG.SIP.DevicePrefix,
		ChannelPrefix: global.CONFIG.SIP.ChannelPrefix,
	}
	global.Logger.Info("sip配置: ", zap.Any("sipServer", sipServer))

	// 初始化媒体服务器
	global.Logger.Info("初始化媒体服务中...")
	var mediaList []system.MediaServer
	result := global.DB.Order("sort ASC").Find(&mediaList)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			global.Logger.Info("数据库中不存在媒体服务器数据，等待从配置文件中加载")
		} else {
			global.Logger.Error("查询媒体服务器错误", zap.Error(result.Error))
		}
	}
	var tmpService system.MediaServer
	for _, server := range mediaList {
		global.MediaServer = &server
		_, err := utils.ZLMHttpRequest(helper.ZlmGetApiList, nil)
		if err != err {
			global.Logger.Error("媒体服务服务器离线", zap.String("mediaServerId", server.MediaServerId))
		} else {
			global.Logger.Info("媒体服务服务器在线", zap.String("mediaServerId", server.MediaServerId))
			if tmpService.MediaServerId == "" || tmpService.Sort > server.Sort {
				tmpService = server
			}
		}
		global.MediaServer = nil
	}
	if tmpService.MediaServerId == "" {
		server := system.MediaServer{}
		global.Logger.Info("未从数据库获取媒体服务器，从配置文件中获取。")
		server.SetAddress(global.CONFIG.Media.Address)
		server.SetRTSPPort(global.CONFIG.Media.RTSPPort)
		server.SetRestful(global.CONFIG.Media.Restful)
		server.SetRTP(global.CONFIG.Media.RTP)
		server.SetRTMPPort(global.CONFIG.Media.RTMPPort)
		server.SetSecret(global.CONFIG.Media.Secret)
		server.SetMediaServerId(global.CONFIG.Media.MediaServerId)
		global.MediaServer = &server
		_, err := utils.ZLMHttpRequest(helper.ZlmGetApiList, nil)
		if err != nil {
			global.Logger.Error("初始化媒体服务器失败")
			return
		}
		server.Status = helper.StreamStart
		global.DB.Create(global.MediaServer)
	} else {
		global.MediaServer = &tmpService
	}
	global.Logger.Info("初始化媒体服务器完成")
	global.Logger.Info("初始化SIP服务中...")
	NewSipServer()
}

func registerHandler() {
	_ = Server.OnRequest(sip.REGISTER, Register)
	_ = Server.OnRequest(sip.MESSAGE, Message)
}
