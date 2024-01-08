package system

// sip实体

import (
	"nebula.xyz/model"
)

type SipServer struct {
	model.NEBULA

	SipId string // SIP服务器ID
	IP    string // SIP服务器地址
	Port  uint   // SIP服务器端口
	Realm string // SIP区域
	// TODO 是否需要加密
	Password      string // SIP密码
	UserAgent     string // SIP用户代理
	DevicePrefix  string // comment:设备前缀
	ChannelPrefix string // 通道前缀
}
