package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"nebula.xyz/global"
	"nebula.xyz/model"
	resp "nebula.xyz/model/response"
	"nebula.xyz/model/system"
	"strconv"
)

type DeviceApi struct{}

// GetAllDeviceInfo 获取所有设备Api
func (d *DeviceApi) GetAllDeviceInfo(c *gin.Context) {
	devices, err := deviceService.GetAllDeviceInfo()
	if err != nil {
		global.Logger.Error("查询设备信息失败")
		model.ErrorWithMessage("获取失败", c)
		return
	}
	model.ErrorWithDetailed(devices, "获取成功", c)
}

// GetDeviceInfoById 根据Id获取设备信息
func (d *DeviceApi) GetDeviceInfoById(c *gin.Context) {
	deviceId := c.Param("deviceId")
	if len(deviceId) != 20 {
		model.ErrorWithMessage("设备Id错误", c)
		return
	}
	device, err := deviceService.GetDeviceInfoById(deviceId)
	if err != nil {
		model.ErrorWithMessage(fmt.Sprintf("获取%s失败", deviceId), c)
		return
	}
	model.OkWithDetailed(device, "获取成功", c)
}

// UpdateDeviceInfo UpdateDeviceInf 根据设备Id更新设备名称，当前只可以更新名称
func (d *DeviceApi) UpdateDeviceInfo(c *gin.Context) {
	var device system.Device
	err := c.ShouldBindJSON(&device)
	if err != nil {
		model.ErrorWithMessage(err.Error(), c)
		return
	}
	if len(device.DeviceId) != 20 || len(device.Name) == 0 {
		model.ErrorWithMessage("参数错误", c)
	}

	err = deviceService.UpdateDeviceInfoById(device)
	if err != nil {
		model.ErrorWithMessage("更新错误", c)
		return
	}
	model.OKWithMessage("更新成功", c)
}

// GenerateDevice 生成设备，以及要生成通道的个数
func (d *DeviceApi) GenerateDevice(c *gin.Context) {
	genChannelNum := c.DefaultQuery("channelNum", "0")
	genInfo := resp.GenerateInfo{}
	device, err := deviceService.GenerateDevice()
	if err != nil {
		model.ErrorWithMessage("生成设备失败", c)
		return
	}
	genInfo.Device = device
	num, _ := strconv.Atoi(genChannelNum)
	if num <= 0 {
		model.OkWithDetailed(genInfo, "生成设备成功", c)
		return
	}
	channels, err := channelService.GenerateChannel(num, device)
	if err != nil {
		model.ErrorWithMessage("生成通道失败", c)
		return
	}
	genInfo.Channels = channels
	genInfo.ChannelSum = len(channels)
	model.OkWithDetailed(genInfo, "生成设备成功", c)
	return
}

// DeleteDevice 删除指定设备
func (d *DeviceApi) DeleteDevice(c *gin.Context) {
	// TODO 删除指定设备
}
