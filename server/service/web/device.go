package web

import (
	"errors"
	"go.uber.org/zap"
	"nebula.xyz/global"
	"nebula.xyz/helper"
	"nebula.xyz/model/request"
	"nebula.xyz/model/system"
	"nebula.xyz/utils"
	"strconv"
)

type DeviceService struct{}

// GetDevicePagination 获取到所有设备信息
func (d *DeviceService) GetDevicePagination(pagination request.Pagination) (devices []system.Device, total int64, err error) {
	db := global.DB.Model(&system.Device{})
	err = db.Count(&total).Error
	if err != nil {
		global.Logger.Error("设备分页查询失败", zap.Error(err))
		return
	}
	if total < 0 {
		return
	}
	offset := (pagination.Page - 1) * pagination.Limit
	err = db.Offset(offset).Limit(pagination.Limit).Find(&devices).Error
	if err != nil {
		global.Logger.Error("设备分页查询失败", zap.Error(err))
	}
	return
}

// GetDeviceInfoById 获取指定设备信息
func (d *DeviceService) GetDeviceInfoById(id string) (device *system.Device, err error) {
	err = global.DB.Where("device_id = ?", id).First(&device).Error
	return
}

// UpdateDeviceInfoById 根据DeviceId 更新设备的名字
func (d *DeviceService) UpdateDeviceInfoById(device system.Device) (err error) {
	err = global.DB.Model(&system.Device{}).Where("device_id = ? ", device.DeviceId).Update("name", &device.Name).Error
	global.Logger.Error("DeviceService", zap.Error(err))
	return err
}

// GenerateDevice 生成一条设备信息
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
	return device, nil
}

func (d *DeviceService) CreateDevice(create request.DeviceCreate) error {
	device := system.Device{}
	tx := global.DB.Begin()
	result := tx.Where("device_id = ?", create.DeviceId).Find(&device)
	if result.RowsAffected != 0 {
		global.Logger.Info("设备id已存在", zap.String("device_id", create.DeviceId))
		tx.Rollback()
		return result.Error
	}
	device.DeviceId = create.DeviceId
	device.Name = create.Name
	device.Realm = global.CONFIG.SIP.Realm
	device.Transport = helper.DefaultTransPort
	device.Port = strconv.Itoa(create.Port)
	device.ChannelCount = 0
	device.Status = helper.DeviceOffline
	createResult := tx.Create(&device)
	if createResult.Error != nil {
		global.Logger.Error("创建设备设备", zap.Error(createResult.Error))
		tx.Rollback()
		return createResult.Error
	}
	tx.Commit()
	return nil
}

// DeleteDevice 删除设备
func (d *DeviceService) DeleteDevice(id string) error {
	result := global.DB.Where("device_id = ?", id).Delete(&system.Device{})
	if result.Error != nil {
		global.Logger.Error("删除设备失败", zap.String("deviceId", id))
		return result.Error
	}
	return nil
}
