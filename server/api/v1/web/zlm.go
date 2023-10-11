package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"nebula.xyz/global"
	"nebula.xyz/helper"
	"nebula.xyz/model/request"
	"nebula.xyz/model/response"
	"nebula.xyz/model/system"
	"nebula.xyz/utils"
	"net/http"
)

type ZlmHookApi struct{}

// OnServerKeepalive ZLM心跳
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

func (z *ZlmHookApi) OnPublish(c *gin.Context) {
	var publish *request.OnPublish
	err := c.ShouldBindJSON(&publish)
	if err != nil {
		c.JSON(http.StatusOK,
			response.ZLMHookResponse{
				Code: helper.ZLMeidaHookFail,
				Msg:  helper.ZLMeidaHookFailMessage,
			})
		return
	}
	global.Logger.Info("收到rtsp/rtmp/rtp推流鉴权事件。", zap.String("schema", publish.Schema))
	// TODO
}

func (z *ZlmHookApi) OnStreamChanged(c *gin.Context) {
	var change *request.StreamChangedData
	err := c.ShouldBindJSON(&change)
	if err != nil {
		c.JSON(http.StatusOK,
			response.ZLMHookResponse{
				Code: helper.ZLMeidaHookFail,
				Msg:  helper.ZLMeidaHookFailMessage,
			})
		return
	}
	if change.Regist != true {
		// 注销
		global.Logger.Info("触发注销操作")

	} else {
		// 注册
		global.Logger.Info("触发注册操作")
		steam := &system.Stream{
			StreamId:         utils.HexToStream(change.Stream),
			VHost:            change.VHost,
			App:              change.App,
			OriginType:       change.OriginTypeStr,
			Schema:           change.Schema,
			TotalReaderCount: change.TotalReaderCount,
		}
		err = zlmService.UpdateStreamInfo(steam)
		if err != nil {
			global.Logger.Info("修改流信息失败")
			c.JSON(http.StatusOK,
				response.ZLMHookResponse{
					Code: helper.ZLMeidaHookFail,
					Msg:  helper.ZLMeidaHookFailMessage,
				})
			return
		}

		global.Logger.Info("更新流信息成功")
		c.JSON(http.StatusOK,
			response.ZLMHookResponse{
				Code: helper.ZLMeidaHookSuccess,
				Msg:  helper.ZLMeidaHookSuccessMessage,
			})
	}
}
