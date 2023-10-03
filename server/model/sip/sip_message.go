package sip

import "encoding/xml"

type SipMessage struct {
	XMLName      xml.Name
	CmdType      string `xml:"CmdType"`
	Sn           string `xml:"SN"`
	DeviceID     string `xml:"DeviceID"`
	DeviceType   string `xml:"DeviceType"`
	Manufacturer string `xml:"Manufacturer"`
	Model        string `xml:"Model"`
	Firmware     string `xml:"Firmware"`
	MaxCamera    int    `xml:"MaxCamera"`
	MaxAlarm     string `xml:"MaxAlarm"`
}
