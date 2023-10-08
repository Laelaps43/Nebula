package sdp

import "fmt"

//	RFC 4566

// Origin defines the structure for the "o=" field
type Origin struct {
	Username       string // 用户名
	SessionId      int    // 识别一个全局对话
	SessionVersion int    // 会话版本
	NetType        string // 网络类型，目前只有IN表示Internet
	AddrType       string // 地址类型 IP4 IP6
	UnicastAddress string // 发送放地址
}

func (o Origin) String() string {
	return fmt.Sprintf(
		"o=%s %d %d %s %s %s",
		o.Username,
		o.SessionId,
		o.SessionVersion,
		o.NetType,
		o.AddrType,
		o.UnicastAddress,
	)
}
