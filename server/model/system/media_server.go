package system

import "nebula.xyz/model"

type MediaServer struct {
	model.NEBULA
	ID       uint   `gorm:"primaryKey comment:主键"`
	RTP      string `gorm:"comment:媒体服务RTP接受端口"`
	Restful  string `gorm:"comment:媒体服务器restful端口"`
	Secret   string `gorm:"comment:媒体服务器secret"`
	Address  string `gorm:"comment:媒体服务器地址address"`
	RTSPPort string `gorm:"comment:媒体体服务器rtsp端口"`
	RTMPPort string `gorm:"comment:流媒体服务器rtmp端口"`
	Status   uint   `gorm:"comment:流媒体服务器状态"`
}
