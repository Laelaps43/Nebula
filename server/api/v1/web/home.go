package web

import (
	"github.com/gin-gonic/gin"
	"nebula.xyz/global"
	"nebula.xyz/model"
)

type HomeApi struct{}

// GetOverView 获取首页设备信息总览
func (h *HomeApi) GetOverView(ctx *gin.Context) {
	overViewResultList := homeService.GetOverView()
	model.OkWithDetailed(overViewResultList, "获取成功", ctx)
	return
}

func (h *HomeApi) GetSystemInfo(ctx *gin.Context) {
	model.OkWithDetailed(global.Info, "获取成功", ctx)
}

func (h *HomeApi) GetServerInfo(ctx *gin.Context) {
	info, err := homeService.GetServerInfo()
	if err != nil {
		model.ErrorWithMessage("获取服务器信息错误", ctx)
		return
	}
	model.OkWithDetailed(info, "获取成功", ctx)
}
