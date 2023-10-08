package sdp_test

import (
	"fmt"
	"nebula.xyz/sip/sdp"
	"testing"
	"time"
)

func TestSdpSession_String(t *testing.T) {
	session := sdp.SdpSession{
		Version: 0,
		Origin: sdp.Origin{
			Username:       "44010200492000000001",
			SessionVersion: 0,
			SessionId:      0,
			NetType:        "IN",
			AddrType:       "IP4",
			UnicastAddress: "192.168.2.11",
		},
		SessionName: "Play",
		ConnectionData: sdp.ConnectionData{
			NetType:           "IN",
			AddrType:          "IP4",
			ConnectionAddress: "192.168.2.11",
		},
		Timing: []sdp.Timing{sdp.Timing{Start: time.Time{}, End: time.Time{}}},
		MediaName: sdp.MediaName{
			Media:   "video",
			Port:    "10000",
			Protos:  []string{"TCP", "RTP", "AVP"},
			Formats: []string{"96", "98", "97"},
		},
		Attributes: []sdp.Attribute{
			{Key: "recvonly"},
			{Key: "setup", Value: "passive"},
			{Key: "connection", Value: "new"},
			{Key: "rtpmap", Value: "96 PS/90000"},
			{Key: "rtpmap", Value: "98 H264/90000"},
			{Key: "rtpmap", Value: "97 MPEG4/90000"},
		},
		SSRC: "0102000001",
	}
	fmt.Println(session.String())
}
