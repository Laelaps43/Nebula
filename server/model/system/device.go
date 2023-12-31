package system

import (
	"go.uber.org/zap"
	"nebula.xyz/global"
	"time"

	"nebula.xyz/model"
)

type Device struct {
	model.NEBULA
	DeviceId      string          `gorm:"primaryKey;comment:设备国标ID" json:"deviceId"` // 设备国标Id
	Realm         string          `gorm:"comment:设备域" json:"realm"`                  // 设备域
	Name          string          `gorm:"comment:设备名称" json:"name"`                  // 设备名称
	DeviceType    string          `gorm:"comment:设备类型" json:"-"`                     // 设备类型
	Firmware      string          `gorm:"comment:设备固件" json:"-"`                     // 设备固件
	Model         string          `gorm:"comment:设备型号" json:"-"`                     // 设备型号
	Manufacturer  string          `gorm:"comment:制造商" json:"manufacturer"`           // 制造商
	Transport     string          `gorm:"comment:传输协议" json:"transport"`             // 传输协议
	StreamModel   string          `gorm:"comment:流传输模式" json:"-"`                    // 流传输模式
	IP            string          `gorm:"comment:设备地址" json:"-"`                     // 设备地址
	Port          string          `gorm:"comment:设备端口" json:"-"`                     // 设备端口
	RegisterAt    *time.Time      `gorm:"comment:设备注册时间" json:"registerAt"`          // 设置注册时间
	KeepLiveAt    *time.Time      `gorm:"comment:心跳时间" json:"keepLiveAt"`            // 心跳时间
	ChannelCount  int             `gorm:"comment:通道个数" json:"channelCount"`          // 通道个数
	Expires       string          `gorm:"comment:有效时间" json:"-"`                     // 有效时间
	Status        string          `gorm:"comment:设备状态：1 在线 0 离线" json:"status"`      // 设备状态
	DeviceChannel []DeviceChannel `gorm:"foreignKey:DeviceId" json:"-"`
}

func (d *Device) TableName() string {
	return "sys_device"
}

// DeviceById 根据Id查询设备信息
func (d *Device) DeviceById() error {
	if err := global.DB.First(&d).Error; err != nil {
		return err
	}
	return nil
}

// DeviceAdd 添加设备
func (d *Device) DeviceAdd() error {
	if err := global.DB.Create(d).Error; err != nil {
		global.Logger.Error("保存设备信息错误！")
		return err
	}
	return nil
}

// DeviceUpdate 更新设备信息
func (d *Device) DeviceUpdate() error {
	if err := global.DB.Model(&Device{}).Where("device_id = ?", d.DeviceId).Updates(d).Error; err != nil {
		global.Logger.Error("更新设备信息失败！")
		return err
	}
	return nil
}

// IsExist 判断设备是否存在
func (d *Device) IsExist() bool {
	if err := d.DeviceById(); err != nil {
		return false
	}
	return true
}

// DeviceCountByStatus 根据状态获取设备数量
func (d *Device) DeviceCountByStatus() int64 {

	var count int64 = 0
	err := global.DB.Model(&Device{}).Where("status = ?", d.Status).Count(&count).Error
	if err != nil {
		global.Logger.Error("获取设备数量失败:状态"+d.Status, zap.Error(err))
	}
	return count
}
