package web

import (
	"go.uber.org/zap"
	"nebula.xyz/global"
	"nebula.xyz/model/system"
)

type DeviceService struct{}

// GetAllDeviceInfo 获取到所有设备信息
func (d *DeviceService) GetAllDeviceInfo() (devices []system.Device, err error) {
	err = global.DB.Find(&devices).Error
	return
}

// GetDeviceInfoById 获取指定设备信息
func (d *DeviceService) GetDeviceInfoById(id string) (device system.Device, err error) {
	err = global.DB.Where("device_id = ?", id).First(&device).Error
	return
}

// UpdateDeviceInfoById 根据DeviceId 更新设备的名字
func (d *DeviceService) UpdateDeviceInfoById(device system.Device) (err error) {
	err = global.DB.Model(&system.Device{}).Where("device_id = ? ", device.DeviceId).Update("name", &device.Name).Error
	global.Logger.Error("DeviceService", zap.Error(err))
	return err
}
