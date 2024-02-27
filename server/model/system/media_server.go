package system

import (
	"nebula.xyz/model"
	"time"
)

type MediaServer struct {
	model.NEBULA
	ID            uint       `gorm:"primaryKey comment:主键"`
	RTP           string     `gorm:"comment:媒体服务RTP接受端口"`
	Restful       string     `gorm:"comment:媒体服务器restful端口"`
	Secret        string     `gorm:"comment:媒体服务器secret"`
	Address       string     `gorm:"comment:媒体服务器地址address"`
	RTSPPort      string     `gorm:"comment:媒体体服务器rtsp端口"`
	RTMPPort      string     `gorm:"comment:流媒体服务器rtmp端口"`
	Status        uint       `gorm:"comment:流媒体服务器状态 0-离线 1-在线 default:0"`
	MediaServerId string     `gorm:"uniqueIndex comment:媒体服务器ID"`
	KeepLiveAt    *time.Time `gorm:"comment:心跳时间"`
	Sort          int        `gorm:"comment:媒体服务器排序"`
	TcpClient     int        `gorm:"comment:Tcp客户端"`
	TcpServer     int        `gorm:"comment:Tcp服务端"`
	TcpSession    int        `gorm:"comment:Tcp会话"`
	UdpServer     int        `gorm:"comment:Udp服务端"`
	UdpSession    int        `gorm:"comment:Udp会话"`
	MediaSource   int        `gorm:"comment:媒体源个数"`
	Domain        string     `gorm:"comment:domain地址"`
}

func (m *MediaServer) TableName() string {
	return "media_server"
}

func (m *MediaServer) GetAddress() string {
	return m.Address
}

func (m *MediaServer) SetAddress(a string) {
	m.Address = a
}

func (m *MediaServer) GetRTP() string {
	return m.RTP
}

func (m *MediaServer) SetRTP(p string) {
	m.RTP = p
}

func (m *MediaServer) GetRestful() string {
	return m.Restful
}

func (m *MediaServer) SetRestful(p string) {
	m.Restful = p
}

func (m *MediaServer) GetRTSPPort() string {
	return m.RTSPPort
}

func (m *MediaServer) SetRTSPPort(p string) {
	m.RTSPPort = p
}

func (m *MediaServer) GetRTMPPort() string {
	return m.RTMPPort
}

func (m *MediaServer) SetRTMPPort(p string) {
	m.RTMPPort = p
}

func (m *MediaServer) GetSecret() string {
	return m.Secret
}

func (m *MediaServer) SetSecret(s string) {
	m.Secret = s
}

func (m *MediaServer) SetMediaServerId(serverId string) {
	m.MediaServerId = serverId
}

func (m *MediaServer) GetMediaServerId() string {
	return m.MediaServerId
}

func (m *MediaServer) GetDomain() string {
	return m.Domain
}
