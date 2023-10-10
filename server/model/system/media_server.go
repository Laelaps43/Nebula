package system

import (
	"nebula.xyz/model"
	"time"
)

type MediaServer struct {
	model.NEBULA
	ID         uint       `gorm:"primaryKey comment:主键"`
	RTP        string     `gorm:"comment:媒体服务RTP接受端口"`
	Restful    string     `gorm:"comment:媒体服务器restful端口"`
	Secret     string     `gorm:"comment:媒体服务器secret"`
	Address    string     `gorm:"comment:媒体服务器地址address"`
	RTSPPort   string     `gorm:"comment:媒体体服务器rtsp端口"`
	RTMPPort   string     `gorm:"comment:流媒体服务器rtmp端口"`
	Status     uint       `gorm:"comment:流媒体服务器状态"`
	KeepLiveAt *time.Time `gorm:"comment:心跳时间"` // 心跳时间
	Sort       int        `gorm:"comment:媒体服务器排序"`
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
