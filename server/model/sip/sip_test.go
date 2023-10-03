package sip

import (
	"fmt"
	"nebula.xyz/helper"
	"testing"
)

func TestQuery(t *testing.T) {
	info := DeviceQueryInfo{DeviceId: "1111", Sn: helper.DeviceOffline, CmdType: "keepalive"}
	fmt.Println(info.ToXML())
}
