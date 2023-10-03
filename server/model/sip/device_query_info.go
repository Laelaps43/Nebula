package sip

import (
	"encoding/xml"
	"nebula.xyz/helper"
)

type DeviceQueryInfo struct {
	XMLName  xml.Name         `xml:"Query"`
	CmdType  helper.QueryType `xml:"CmdType"`
	Sn       string           `xml:"SN"`
	DeviceId string           `xml:"DeviceID"`
}

func (info *DeviceQueryInfo) ToXML() string {
	x, _ := xml.MarshalIndent(info, " ", "  ")
	return xml.Header + string(x)
}
