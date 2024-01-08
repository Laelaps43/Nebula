package web

import (
	"github.com/gin-gonic/gin"
	"nebula.xyz/api/v1/web"
)

type RecordRouter struct{}

func (h *RecordRouter) InitRecordRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	recordRouter := Router.Group("record")
	recordApi := web.WebApiAll.RecordApi
	{
		recordRouter.POST("page", recordApi.GetAllVideoRecord)
		recordRouter.POST("range", recordApi.GetVideoDateRange)
		recordRouter.GET("select/:stream/:date", recordApi.GetSelectRecord)
		recordRouter.GET("play/:stream/:id", recordApi.GetRecordPlay)
	}
	return recordRouter
}
