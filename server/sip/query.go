package sip

import (
	"github.com/ghettovoice/gosip/sip"
	"go.uber.org/zap"
	"nebula.xyz/global"
	"nebula.xyz/helper"
	sb "nebula.xyz/model/sip"
	"nebula.xyz/model/system"
	"nebula.xyz/utils"
)

func QueryDeviceSip(device system.Device) {
	body := sb.DeviceQueryInfo{DeviceId: device.DeviceId, CmdType: helper.DeviceInfoCmdType, Sn: utils.GetSn()}

	request, _ := createMessageRequest(device, contentTypeXML, sip.MESSAGE, body.ToXML())
	global.Logger.Debug("生成的查询设备信息XML", zap.String("XML", body.ToXML()))
	tx, err := Server.Request(request)
	if err != nil {
		global.Logger.Error("查询设备出错", zap.Error(err))
		return
	}
	resp := tx.Responses()
	global.Logger.Info("收到Invite通知", zap.Any("response", resp))
}
