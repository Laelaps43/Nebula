package system

import (
	"go.uber.org/zap"
	"nebula.xyz/global"
	"nebula.xyz/model"
)

type DeviceChannel struct {
	model.NEBULA
	ChannelId    string `gorm:"primaryKey;comment:通道ID" xml:"DeviceID" json:"channelId"` // 通道ID
	DeviceId     string `gorm:"comment:设备ID" json:"-"`                                   // 设备Id
	Name         string `gorm:"comment:设备名称" xml:"Name" json:"name"`                     // 设备名称
	Manufacturer string `gorm:"comment:制造商" xml:"Manufacturer" json:"manufacturer"`      // 制造商
	Model        string `gorm:"comment:平台型号" xml:"Model" json:"model"`                   // 平台型号
	Transport    string `gorm:"comment:传输协议" json:"transport"`                           // 传输协议
	Owner        string `gorm:"comment:平台归属" xml:"Owner" json:"-"`                       // 平台归属
	CivilCode    string `gorm:"comment:行政区域" xml:"CivilCode" json:"-"`                   // 行政区域
	Address      string `gorm:"comment:平台安装地址" xml:"Address" json:"address"`             // 安装地址
	Parental     string `gorm:"comment:是否有子设备，1有，0没有" xml:"Parental" json:"-"`           // 当为设备时，是否有子设备，1有，0没有
	ParentID     string `gorm:"comment:父设备/区域/系统ID" xml:"ParentID" json:"-"`             // 父设备/区域/系统ID
	SafetyWay    string `gorm:"comment:信令安全模式" xml:"SafetyWay" json:"-"`                 //信令安全模式，0不采用、2 S/MIME签名方式、3 S/MIME加密他签名同时采用方式、4 数字摘要方式
	RegisterWay  string `gorm:"comment:注册方式" xml:"RegisterWay" json:"-"`                 // 注册方式，1 标准认证注册模式 、2 基于口令的双向认证模式、3 基于数字证书的双向认证注册模式
	Secrecy      string `gorm:"comment:保密属性" xml:"Secrecy" json:"-"`                     // 保密属性，0不涉密、1涉密
	Status       string `gorm:"comment:通道状态" xml:"Status" json:"status"`                 // 通道状态
}

func (d *DeviceChannel) TableName() string {
	return "device_channel"
}

// DeviceChannelById 根据通道Id查找通道
func (d *DeviceChannel) DeviceChannelById() error {
	if err := global.DB.Where("channel_id = ?", d.ChannelId).First(&d).Error; err != nil {
		return err
	}
	return nil
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

// IsExist 判断通道是否存在
func (d *DeviceChannel) IsExist() bool {
	err := d.DeviceChannelById()
	if err != nil || d.ChannelId == "" {
		return false
	}
	return true
}

// GetAllChannels 获取所有通道
func (d *DeviceChannel) GetAllChannels() (channels []DeviceChannel, err error) {
	if err = global.DB.Model(&DeviceChannel{}).Find(&channels).Error; err != nil {
		return nil, err
	}
	return
}

// UpdateChannelInfoById 根据Id更新通道信息
func (d *DeviceChannel) UpdateChannelInfoById() (err error) {
	if err = global.DB.Model(&DeviceChannel{}).Where("channel_id = ?", d.ChannelId).Updates(d).Error; err != nil {
		return err
	}
	return
}

// DeleteChannelById 根据Id删除通道
func (d *DeviceChannel) DeleteChannelById() (err error) {
	err = global.DB.Delete(d).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *DeviceChannel) ChannelCount() int64 {
	var count int64 = 0
	err := global.DB.Model(&DeviceChannel{}).Count(&count).Error
	if err != nil {
		global.Logger.Error("获取通道数量失败", zap.Error(err))
	}
	return count
}
