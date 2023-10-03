package system

import (
	"nebula.xyz/global"
	"nebula.xyz/model"
)

type DeviceChannel struct {
	model.NEBULA
	ChannelId    string `gorm:"primaryKey;comment:通道ID" xml:"DeviceID"` // 通道ID
	DeviceId     string `gorm:"comment:设备ID"`                           // 设备Id
	Name         string `gorm:"comment:设备名称" xml:"Name"`                // 设备名称
	Manufacturer string `gorm:"comment:制造商" xml:"Manufacturer"`         // 制造商
	Model        string `gorm:"comment:平台型号" xml:"Model"`               // 平台型号
	Transport    string `gorm:"comment:传输协议" `                          // 传输协议
	Owner        string `gorm:"comment:平台归属" xml:"Owner"`               // 平台归属
	CivilCode    string `gorm:"comment:行政区域" xml:"CivilCode"`           // 行政区域
	Address      string `gorm:"comment:平台安装地址" xml:"Address"`           // 安装地址
	Parental     string `gorm:"comment:是否有子设备，1有，0没有" xml:"Parental"`   // 当为设备时，是否有子设备，1有，0没有
	ParentID     string `gorm:"comment:父设备/区域/系统ID" xml:"ParentID"`     // 父设备/区域/系统ID
	SafetyWay    string `gorm:"comment:信令安全模式" xml:"SafetyWay"`         //信令安全模式，0不采用、2 S/MIME签名方式、3 S/MIME加密他签名同时采用方式、4 数字摘要方式
	RegisterWay  string `gorm:"comment:注册方式" xml:"RegisterWay"`         // 注册方式，1 标准认证注册模式 、2 基于口令的双向认证模式、3 基于数字证书的双向认证注册模式
	Secrecy      string `gorm:"comment:保密属性" xml:"Secrecy"`             // 保密属性，0不涉密、1涉密
	Status       string `gorm:"comment:通道状态" xml:"Status"`              // 通道状态
}

func (d *DeviceChannel) TableName() string {
	return "device_channel"
}

// DeviceChannelById 根据通道Id查找通道
func (d *DeviceChannel) DeviceChannelById() (DeviceChannel, error) {
	tmp := DeviceChannel{}
	if err := global.DB.Where("channel_id = ?", d.ChannelId).First(&tmp).Error; err != nil {
		return DeviceChannel{}, err
	}
	return tmp, nil
}

// ChannelUpdate 更新通道信息
func (d *DeviceChannel) ChannelUpdate() error {
	if err := global.DB.Where("channel_id = ?", d.ChannelId).Updates(d).Error; err != nil {
		return err
	} else {
		return nil
	}
}

// ChannelAdd 添加通道
func (d *DeviceChannel) ChannelAdd() error {
	if err := global.DB.Create(d).Error; err != nil {
		return err
	} else {
		return nil
	}
}
