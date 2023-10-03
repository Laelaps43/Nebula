package sip

// GB28181系统目录（通道查询）

import (
	"encoding/xml"
	"nebula.xyz/helper"
)

type DeviceQueryCatalog struct {
	XMLName  xml.Name         `xml:"Query"`
	CmdType  helper.QueryType `xml:"CmdType"`
	SN       string           `xml:"SN"`
	DeviceID string           `xml:"DeviceID"`
}
