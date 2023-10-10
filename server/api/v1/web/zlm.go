package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"nebula.xyz/global"
	"nebula.xyz/helper"
	"nebula.xyz/model/request"
	"nebula.xyz/model/response"
	"nebula.xyz/model/system"
	"net/http"
)

type ZlmHookApi struct{}

func (z *ZlmHookApi) OnServerKeepalive(c *gin.Context) {
	var keepalive request.ServerKeepalive
	if err := c.ShouldBindJSON(&keepalive); err != nil {
		global.Logger.Info("获取ZLM数据失败")
		c.JSON(http.StatusOK,
			response.ZLMHookResponse{
				Code: helper.ZLMeidaHookFail,
				Msg:  helper.ZLMeidaHookFailMessage,
			})
		return
	}
	global.Logger.Info(fmt.Sprintf("收到ZLM id %s 的心跳", keepalive.MediaServerId))
	mediaInfo := &system.MediaServer{}
	mediaInfo.Address = keepalive.MediaServerId
	zlmService.UpdateServerStatus(mediaInfo)
	c.JSON(http.StatusOK,
		response.ZLMHookResponse{
			Code: helper.ZLMeidaHookSuccess,
			Msg:  helper.ZLMeidaHookSuccessMessage,
		})
	return
}
