package web

import (
	"github.com/gin-gonic/gin"
	"nebula.xyz/model"
)

type HomeApi struct{}

// GetOverView 获取首页设备信息总览
func (h *HomeApi) GetOverView(ctx *gin.Context) {

	overViewResultList := homeService.GetOverView()
	model.OkWithDetailed(overViewResultList, "获取成功", ctx)
	return
}
