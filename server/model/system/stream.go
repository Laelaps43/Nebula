package system

import (
	"nebula.xyz/global"
	"nebula.xyz/model"
	"time"
)

type Stream struct {
	model.NEBULA
	ChannelId        string `gorm:"comment:通道ID"`
	DeviceId         string `gorm:"comment:设备ID"`
	StreamType       string `gorm:"comment:pull服务器拉流，push设备推流"`
	Status           uint   `gorm:"comment:流状态 1-关闭 2-正常"`
	StreamId         string `gorm:"primaryKey, comment:流ID SSRC"`
	HTTP             string `gorm:"comment:mu38地址"`
	RTMP             string `gorm:"comment:RTMP地址"`
	RTSP             string `gorm:"comment:RTSP地址"`
	WSFLV            string `gorm:"comment:WSFLV地址"`
	ZlmAccept        bool   `gorm:"comment:zlm是否收到流"`
	App              string `gorm:"comment:zlm应用名"`
	OriginType       int    `gorm:"comment:zlm产生流类型"`
	Schema           string `gorm:"comment:zlm流协议"`
	TotalReaderCount int    `gorm:"comment:zlm观看总人数"`
	VHost            string `gorm:"comment:zlm虚拟主机"`
	Record           uint   `gorm:"comment:流是否被录制 1-否 2-是"`

	Start, End time.Time `gorm:"-"`
}

func (s *Stream) TableName() string {
	return "stream"
}

// GetStreamById 根据Stream Id 查找Stream
func (s *Stream) GetStreamById() error {
	if err := global.DB.Where("stream_id = ?", s.StreamId).First(&s).Error; err != nil {
		return err
	}
	return nil
}

// Save 保存流信息
func (s *Stream) Save() error {
	if err := global.DB.Create(s).Error; err != nil {
		global.Logger.Error("保存流信息错误")
		return err
	}
	return nil
}

func (s *Stream) Update() error {
	if err := global.DB.Where("stream_id = ?", s.StreamId).Updates(s).Error; err != nil {
		return err
	}
	return nil
}
func (s *Stream) GetStreamByDeviceAndChannel() error {
	if err := global.DB.Where("device_id = ? and channel_id = ?", s.DeviceId, s.ChannelId).First(&s).Error; err != nil {
		return err
	}
	return nil
}
