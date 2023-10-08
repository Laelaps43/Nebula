package sdp

// RFC 4566 Connection Data
// c=<nettype> <addrtype> <connection-address>

import "fmt"

type ConnectionData struct {
	NetType           string // 网络标识，目前之后 IN
	AddrType          string // 地址类型 IP4 IP6
	ConnectionAddress string // 连接地址
}

func (c *ConnectionData) String() string {
	return fmt.Sprintf("c=%s %s %s", c.NetType, c.AddrType, c.ConnectionAddress)
}

func (c *ConnectionData) IsNil() bool {
	if len(c.NetType) == 0 || len(c.AddrType) == 0 || len(c.ConnectionAddress) == 0 {
		return true
	}
	return false
}
