package web

import (
	"errors"
	"go.uber.org/zap"
	"nebula.xyz/global"
	"nebula.xyz/helper"
	"nebula.xyz/model/system"
	"nebula.xyz/utils"
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

// GenerateDevice 生成设备信息
func (d *DeviceService) GenerateDevice() (device *system.Device, err error) {
	device = &system.Device{}
	server := system.SipServer{}
	err = server.GetSipServerOnLine()
	if err != nil {
		global.Logger.Error("获取SipServer信息错误", zap.Error(err))
		return nil, errors.New("获取SipServer信息错误")
	}
	// 生成设备ID
	for {
		randInt := utils.RandInt(6)
		key := server.DevicePrefix + randInt
		device.DeviceId = key
		exist := device.IsExist()
		if !exist {
			break
		}
	}
	device.Status = helper.DeviceOffline
	device.Realm = server.Realm
	err = device.DeviceAdd()
	if err != nil {
		global.Logger.Error("添加设备失败")
		return nil, err
	}
	return device, nil
}
