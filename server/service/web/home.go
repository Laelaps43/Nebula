package web

import (
	"go.uber.org/zap"
	"nebula.xyz/global"
	"nebula.xyz/helper"
	resp "nebula.xyz/model/response"
	"nebula.xyz/model/system"
	"strconv"
	"time"
)

type HomeService struct{}

// GetOverView 获取首页数据总览
func (s HomeService) GetOverView() resp.OverViewResult {

	device := system.Device{}
	// 获取在线设备
	device.Status = helper.DeviceOnline
	onLineCount := device.DeviceCountByStatus()
	// 获取离线设备
	device.Status = helper.DeviceOffline
	offLineCount := device.DeviceCountByStatus()

	// 获取通道
	channel := system.DeviceChannel{}
	channelCount := channel.ChannelCount()

	// 获取正在录像设备
	var videoCount int64 = 0
	err := global.DB.Model(&system.Stream{}).Where("record = ?", helper.StreamRecorded).Count(&videoCount).Error
	if err != nil {
		global.Logger.Error("获取录像设备数量错误", zap.Error(err))
		videoCount = 0
	}
	result := resp.OverViewResult{
		OnlineDevice:  onLineCount,
		OfflineDevice: offLineCount,
		Channel:       channelCount,
		Video:         videoCount,
	}
	return result
}

func (s HomeService) GetServerInfo() (resp.ServerInfo, error) {
	var mediaServer system.MediaServer
	err := global.DB.Model(&system.MediaServer{}).
		Where("media_server_id = ? ", global.MediaServer.GetMediaServerId()).
		First(&mediaServer).Error
	if err != nil {
		global.Logger.Error("获取媒体服务器状态失败", zap.Error(err))
		return resp.ServerInfo{}, err
	}
	mediaServerInfo := resp.MediaServerDetails{
		MediaServiceAddress: global.MediaServer.GetAddress(),
		MediaUniqueID:       global.MediaServer.GetMediaServerId(),
		RTPPort:             global.MediaServer.GetRTP(),
		RestfulPort:         global.MediaServer.GetRestful(),
		RTSPPort:            global.MediaServer.GetRTSPPort(),
		RTMPPort:            global.MediaServer.GetRTMPPort(),
		TCPSessions:         mediaServer.TcpSession,
		UDPSessions:         mediaServer.UdpSession,
		LastHeartbeatTime:   (*mediaServer.KeepLiveAt).Format("2006-01-02 15:04:01"),
	}
	server := resp.ServerDetails{
		ServiceAddress:  global.CONFIG.SIP.Ip,
		SIPServerID:     global.CONFIG.SIP.SipId,
		SIPServerDomain: global.CONFIG.SIP.Realm,
		SIPPassword:     global.CONFIG.SIP.Password,
		Uptime:          getServerInfoUpTime(),
	}
	return resp.ServerInfo{
		ServerDetails:      server,
		MediaServerDetails: mediaServerInfo,
	}, nil
}

func getServerInfoUpTime() int64 {
	serverStartTime, err := global.CACHE.Get(helper.CacheServerUpTimeKey)
	if err != nil {
		return 0
	}
	startTime, _ := strconv.ParseInt(serverStartTime.(string), 10, 64)
	return time.Now().UnixMilli() - startTime
}
