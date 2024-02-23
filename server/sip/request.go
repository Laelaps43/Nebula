package sip

import (
	"fmt"
	"github.com/ghettovoice/gosip/sip"
	"nebula.xyz/global"
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
	to := newToAddress(&device)
	global.Logger.Debug("创建请求，响应地址为" + to.Uri.Host())
	requestBuilder.SetTo(to)
	requestBuilder.SetRecipient(to.Uri)
	requestBuilder.AddVia(newVia(&device))
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

func createVideoRequest(
	channel *system.DeviceChannel,
	device *system.Device,
	method sip.RequestMethod,
	body string,
) (sip.Request, *sip.RequestBuilder) {
	rb := sip.NewRequestBuilder()
	rb.SetFrom(newFromAddress(newParams(map[string]string{"tag": utils.RandString(32)})))
	to := newChannelAddress(channel, device)
	rb.SetTo(to)
	rb.SetRecipient(to.Uri)
	rb.AddVia(newVia(device))
	contentType := sip.ContentType(contentTypeSDP)
	rb.SetContentType(&contentType)
	rb.SetMethod(method)
	userAgent := sip.UserAgentHeader(sipServer.UserAgent)
	rb.SetUserAgent(&userAgent)
	rb.AddHeader(&sip.GenericHeader{
		HeaderName: "Subject",
		Contents:   fmt.Sprintf("%s:%s,%s:%s", channel.ChannelId, device.Realm, device.DeviceId, device.Realm),
	})
	rb.SetBody(body)
	// TODO 序列号
	rb.SetSeqNo(1)
	request, _ := rb.Build()
	return request, rb
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
func newToAddress(device *system.Device) *sip.Address {
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
func newVia(device *system.Device) *sip.ViaHop {
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

// 返回通道地址
func newChannelAddress(channel *system.DeviceChannel, device *system.Device) *sip.Address {
	port64, _ := strconv.Atoi(device.Port)
	port := sip.Port(port64)
	return &sip.Address{
		Uri: &sip.SipUri{
			FUser: sip.String{Str: channel.ChannelId},
			FHost: device.IP,
			FPort: &port,
		},
	}
}
