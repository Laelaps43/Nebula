package web

import (
	"github.com/gin-gonic/gin"
	"nebula.xyz/global"
	"nebula.xyz/model"
	"nebula.xyz/model/system"
	"nebula.xyz/sip"
)

type VideoApi struct{}

func (h *VideoApi) PlayVideo(ctx *gin.Context) {
	
	global.Logger.Info("执行PlayVideo函数")
	sip.Play(&system.Stream{
		ChannelId: "37070000081318000012",
		DeviceId:  "37070000081118000001",
	})
	model.OKWithMessage("Hello!!!", ctx)
}
