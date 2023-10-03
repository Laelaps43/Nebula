package sip

import (
	"fmt"
	"github.com/ghettovoice/gosip/sip"
	"nebula.xyz/model/system"
	"nebula.xyz/utils"
	"strconv"
	"time"
)

const (
	contentTypeXML = "Application/MANSCDP+xml"
	contentTypeSDP = "APPLICATION/SDP"
)

func createMessageRequest(
	device system.Device,
	contentType string,
	method sip.RequestMethod,
	body string) (sip.Request, *sip.RequestBuilder) {
	requestBuilder := sip.NewRequestBuilder()
	requestBuilder.SetFrom(newFromAddress(newParams(map[string]string{"tag": utils.RandString(32)})))
	to := newToAddress(device)
	requestBuilder.SetTo(to)
	requestBuilder.SetRecipient(to.Uri)
	requestBuilder.AddVia(newVia(device))
	requestContentType := sip.ContentType(contentType)
	requestBuilder.SetContentType(&requestContentType)
	requestBuilder.SetMethod(method)
	userAgent := sip.UserAgentHeader(sipServer.UserAgent)
	requestBuilder.SetUserAgent(&userAgent)
	requestBuilder.SetBody(body)
	// TODO 序列号
	//requestBuilder.SetSeqNo()
	request, _ := requestBuilder.Build()
	return request, requestBuilder
}

// 处理参数
func newParams(m map[string]string) *sip.Params {
	params := sip.NewParams()
	for k, v := range m {
		params.Add(k, sip.String{Str: v})
	}
	return &params
}

// 本地地址
func newFromAddress(params *sip.Params) *sip.Address {
	port := sip.Port(sipServer.Port)
	return &sip.Address{
		Uri: &sip.SipUri{
			FUser: sip.String{Str: sipServer.SipId},
			FHost: sipServer.Realm,
			FPort: &port,
		},
		Params: *params,
	}
}

// 返回地址
func newToAddress(device system.Device) *sip.Address {
	port64, _ := strconv.Atoi(device.Port)
	port := sip.Port(port64)
	return &sip.Address{
		Uri: &sip.SipUri{
			FUser: sip.String{Str: device.DeviceId},
			FHost: device.IP,
			FPort: &port,
		},
	}
}

// via
func newVia(device system.Device) *sip.ViaHop {
	port64, _ := strconv.Atoi(device.Port)
	port := sip.Port(port64)
	params := newParams(map[string]string{
		"branch": fmt.Sprintf("z9hG4bK%d", time.Now().UnixMilli()),
	})
	return &sip.ViaHop{
		ProtocolName:    "SIP",
		ProtocolVersion: "2.0",
		Transport:       device.Transport,
		Host:            sipServer.SipId,
		Port:            &port,
		Params:          *params,
	}
}
