package sip_test

import (
	"fmt"
	"nebula.xyz/model/system"
	"nebula.xyz/sip"
	"testing"
)

func TestGetSSRC(t *testing.T) {
	c := &system.DeviceChannel{ChannelId: "37070000081318000012"}
	s := sip.GetSSRC(c)
	fmt.Println(s)
}
