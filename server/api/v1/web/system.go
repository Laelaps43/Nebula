package web

import (
	"github.com/gin-gonic/gin"
	"nebula.xyz/global"
	"nebula.xyz/model"
)

type SystemApi struct{}

func (s *SystemApi) GetSystemInfo(ctx *gin.Context) {
	model.OkWithDetailed(global.Info, "获取成功", ctx)
}
