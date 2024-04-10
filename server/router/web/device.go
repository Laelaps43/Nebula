package web

import (
	"github.com/gin-gonic/gin"
	"nebula.xyz/api/v1/web"
)

type DeviceRouter struct{}

// InitDeviceRouter 添加路由
func (d *DeviceRouter) InitDeviceRouter(group *gin.RouterGroup) {
	deviceRouter := group.Group("device")

	deviceApi := web.WebApiAll.DeviceApi
	{
		// 获取所有Device信息
		deviceRouter.POST("list", deviceApi.GetDeviceInfoPagination)
		deviceRouter.POST("update", deviceApi.UpdateDeviceInfo)
		deviceRouter.GET("create/generate", deviceApi.GenerateDevice)
		deviceRouter.POST("create/create", deviceApi.CreateDevice)
		deviceRouter.GET("delete/:deviceId", deviceApi.DeleteDevice)
	}
}
