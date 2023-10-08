package system

import (
	"nebula.xyz/global"
	"nebula.xyz/model"
	"time"
)

type Stream struct {
	model.NEBULA
	ChannelId  string `gorm:"comment:通道ID"`
	DeviceId   string `gorm:"comment:设备ID"`
	StreamType string `gorm:"comment:pull服务器拉流，push设备推流"`
	Status     string `gorm:"comment:流状态 0 正常 1 关闭 -1 尚未开始"`
	StreamId   string `gorm:"primaryKey, comment:流ID SSRC"`
	HTTP       string `gorm:"comment:mu38地址"`
	RTMP       string `gorm:"comment:RTMP地址"`
	RTSP       string `gorm:"comment:RTSP地址"`
	WSFLV      string `gorm:"comment:WSFLV地址"`
	ZlmAccept  bool   `gorm:"comment:zlm是否收到流"`

	Start, End time.Time `gorm:"-"`
}

// GetStreamById 根据Stream Id 查找Stream
func (s *Stream) GetStreamById() (*Stream, error) {
	tmp := &Stream{}
	if err := global.DB.Where("stream_id = ?", s.StreamId).First(&tmp).Error; err != nil {
		return &Stream{}, err
	}
	return tmp, nil
}

// Save 保存流信息
func (s *Stream) Save() error {
	if err := global.DB.Create(s).Error; err != nil {
		global.Logger.Error("保存流信息错误")
		return err
	}
	return nil
}
