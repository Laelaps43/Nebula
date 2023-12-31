package zlm

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"nebula.xyz/global"
	"nebula.xyz/helper"
	"nebula.xyz/model/request/zlm"
	"nebula.xyz/model/response"
	"net/http"
)

type ZlmHookApi struct{}

// OnServerKeepalive ZLM心跳
func (z *ZlmHookApi) OnServerKeepalive(c *gin.Context) {
	var keepalive zlm.ServerKeepalive
	if err := c.ShouldBindJSON(&keepalive); err != nil {
		global.Logger.Info("绑定ZLM KeepAlive失败", zap.Error(err))
		c.JSON(http.StatusOK,
			response.ZLMHookResponse{
				Code: helper.ZLMeidaHookFail,
				Msg:  helper.ZLMeidaHookFailMessage,
			})
		return
	}
	global.Logger.Info(fmt.Sprintf("收到ZLM id %s 的心跳", keepalive.MediaServerId))
	zlmService.UpdateServerStatus(keepalive)
	c.JSON(http.StatusOK,
		response.ZLMHookResponse{
			Code: helper.ZLMeidaHookSuccess,
			Msg:  helper.ZLMeidaHookSuccessMessage,
		})
	return
}

func (z *ZlmHookApi) OnPublish(c *gin.Context) {
	var publish *zlm.OnPublish
	err := c.ShouldBindJSON(&publish)
	if err != nil {
		global.Logger.Error("OnPublish", zap.Error(err))
		c.JSON(http.StatusOK,
			response.ZLMHookResponse{
				Code: helper.ZLMeidaHookFail,
				Msg:  helper.ZLMeidaHookFailMessage,
			})
		return
	}
	global.Logger.Info("收到rtsp/rtmp/rtp推流鉴权事件。", zap.String("schema", publish.Schema))
	// TODO
	push := response.OnPublish{
		Code: helper.ZLMeidaHookSuccess,
		Msg:  helper.ZLMeidaHookSuccessMessage,
	}
	marshal, err := json.Marshal(push)
	bytes, err := json.Marshal(publish)
	global.Logger.Info("推流鉴权", zap.String("json", string(bytes)))
	global.Logger.Info("推流鉴权", zap.String("json", string(marshal)))
	c.JSON(http.StatusOK, push)
	return
}

func (z *ZlmHookApi) OnStreamChanged(c *gin.Context) {
	var streamChange zlm.StreamChange
	err := c.ShouldBindJSON(&streamChange)
	if err != nil {
		global.Logger.Error("ZLM 流改变参数错误：", zap.Error(err))
		c.JSON(http.StatusOK,
			response.ZLMHookResponse{
				Code: helper.ZLMeidaHookFail,
				Msg:  helper.ZLMeidaHookFailMessage,
			})
		return
	}

}

func (z *ZlmHookApi) OnPlay(ctx *gin.Context) {

}
