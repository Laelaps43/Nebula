package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"nebula.xyz/global"
	"nebula.xyz/model"
	"nebula.xyz/model/request"
	resp "nebula.xyz/model/response"
	"nebula.xyz/model/system"
)

type DeviceApi struct{}

// GetDeviceInfoPagination 分页获取所有设备Api
func (d *DeviceApi) GetDeviceInfoPagination(c *gin.Context) {
	var pagination request.Pagination

	err := c.ShouldBindJSON(&pagination)
	if err != nil {
		model.ErrorWithMessage("请检查分页数据", c)
		return
	}
	devicePagination, total, err := deviceService.GetDevicePagination(pagination)
	if err != nil {
		global.Logger.Error("查询设备信息失败")
		model.ErrorWithMessage("获取失败", c)
		return
	}
	model.OkWithDetailed(resp.PaginationResult{
		List:  devicePagination,
		Total: total,
	}, "获取成功", c)
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

// GenerateDevice 生成临时一个设备
func (d *DeviceApi) GenerateDevice(c *gin.Context) {
	device, err := deviceService.GenerateDevice()
	if err != nil {
		model.ErrorWithMessage("生成设备失败", c)
		return
	}
	model.OkWithDetailed(device.DeviceId, "生成设备成功", c)
	return
}

// CreateDevice 创建设备
func (d *DeviceApi) CreateDevice(c *gin.Context) {
	var deviceCreate request.DeviceCreate

	err := c.ShouldBindJSON(&deviceCreate)
	if err != nil {
		model.ErrorWithMessage("服务器内部异常", c)
		c.Abort()
	}

	err = deviceService.CreateDevice(deviceCreate)
	if err != nil {
		global.Logger.Error("创建设备失败", zap.Error(err))
		model.ErrorWithMessage("创建设备失败", c)
		c.Abort()
	}
	model.OKWithMessage("创建设备成功", c)
}

// DeleteDevice 删除设备
func (d *DeviceApi) DeleteDevice(c *gin.Context) {
	deviceId := c.Param("deviceId")

	if len(deviceId) != 20 {
		model.ErrorWithMessage("设备Id错误", c)
		return
	}
	err := deviceService.DeleteDevice(deviceId)
	if err != nil {
		model.ErrorWithMessage("删除设备失败", c)
		return
	}
	model.OKWithMessage("删除设备成功", c)
}
