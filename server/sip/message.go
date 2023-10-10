package sip

import (
	"encoding/xml"
	"errors"
	"github.com/ghettovoice/gosip/sip"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"nebula.xyz/global"
	sb "nebula.xyz/model/sip"
	"nebula.xyz/model/system"
	"net/http"
	"time"
)

func Message(request sip.Request, transaction sip.ServerTransaction) {

	//global.Logger.Error("MESSAGE-Request:", zap.Any("request", request))
	if l, ok := request.ContentLength(); l.Equals(0) || !ok {
		global.Logger.Info("请求的消息体长度为0")
		_ = transaction.Respond(sip.NewResponseFromRequest("", request, http.StatusOK, http.StatusText(http.StatusOK), ""))
		return
	}

	body := request.Body()

	message := &sb.SipMessage{}

	if err := xml.Unmarshal([]byte(body), message); err != nil {
		global.Logger.Error("解析MESSAGE-Body失败", zap.Error(err))
		return
	}
	if message.XMLName.Local == "Response" {
		switch message.CmdType {
		case "DeviceInfo":
			// 查询设备信息
			device := &system.Device{
				Manufacturer: message.Manufacturer,
				Model:        message.Model,
				DeviceType:   message.DeviceType,
				Firmware:     message.Firmware,
				DeviceId:     message.DeviceID,
				ChannelCount: message.MaxCamera,
			}
			_ = device.DeviceUpdate()
		case "Catalog":
			// TODO 媒体状态处理
			for _, item := range message.DeviceList {
				err := item.DeviceChannelById()
				channel := item
				item.DeviceId = message.DeviceID
				item.ParentID = message.DeviceID
				via, _ := request.ViaHop()
				item.Transport = via.Transport
				global.Logger.Info("添加通道", zap.String("id", channel.ChannelId))
				if !errors.Is(err, gorm.ErrRecordNotFound) {
					_ = item.ChannelUpdate()
				} else {
					_ = item.ChannelAdd()
				}
			}
		}
		_ = transaction.Respond(sip.NewResponseFromRequest("", request, http.StatusOK, http.StatusText(http.StatusOK), ""))
	} else if message.XMLName.Local == "Notify" {
		nowTime := time.Now()
		switch message.CmdType {
		case "Keepalive":
			// 心跳
			device := &system.Device{
				DeviceId:   message.DeviceID,
				KeepLiveAt: &nowTime,
			}
			_ = device.DeviceUpdate()
			err := transaction.Respond(sip.NewResponseFromRequest("", request, http.StatusOK, http.StatusText(http.StatusOK), ""))
			if err != nil {
				global.Logger.Error("回复Keepalive错误", zap.Error(err))
				// 再次回复
				_ = transaction.Respond(sip.NewResponseFromRequest("", request, http.StatusOK, http.StatusText(http.StatusOK), ""))
				return
			}
		}
	}
}
