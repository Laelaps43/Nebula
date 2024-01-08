package web

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"nebula.xyz/global"
	"nebula.xyz/model"
	"nebula.xyz/model/request"
	resp "nebula.xyz/model/response"
	"time"
)

type RecordApi struct{}

func (r *RecordApi) GetAllVideoRecord(ctx *gin.Context) {
	var pagination request.Pagination

	err := ctx.ShouldBindJSON(&pagination)
	if err != nil {
		model.ErrorWithMessage("请检查分页数据", ctx)
		return
	}
	recordPagination, total, err := recordService.GetAllVideoRecord(pagination)
	if err != nil {
		global.Logger.Error("查询录像信息失败")
		model.ErrorWithMessage("获取失败", ctx)
		return
	}
	model.OkWithDetailed(resp.PaginationResult{
		List:  recordPagination,
		Total: total,
	}, "获取成功", ctx)
}

func (r *RecordApi) GetVideoDateRange(ctx *gin.Context) {
	var rangeTime request.RecordRange

	err := ctx.ShouldBindJSON(&rangeTime)
	if err != nil {
		global.Logger.Error("绑定数据错误", zap.Error(err))
		model.ErrorWithMessage("绑定数据错误", ctx)
		return
	}
	recordList := recordService.GetVideoDateRange(rangeTime)

	model.OkWithDetailed(recordList, "获取成功", ctx)
}

func (r *RecordApi) GetSelectRecord(ctx *gin.Context) {
	var date string
	var stream string

	date = ctx.Param("date")
	stream = ctx.Param("stream")
	selectTime, err := time.Parse("2006-01-02", date)
	if err != nil {
		global.Logger.Debug("处理时间错误")
		model.ErrorWithMessage("参数错误", ctx)
		return
	}
	recordList, err := recordService.GetSelectRecord(stream, selectTime)
	if err != nil {
		model.ErrorWithMessage("获取记录失败", ctx)
		return
	}
	model.OkWithDetailed(recordList, "获取成功", ctx)
}

func (r *RecordApi) GetRecordPlay(ctx *gin.Context) {
	var id string
	var stream string

	id = ctx.Param("id")
	stream = ctx.Param("stream")
	url, err := recordService.GetRecordPlay(id, stream)
	if err != nil {
		model.ErrorWithMessage("播放失败", ctx)
		return
	}
	model.OkWithDetailed(url, "获取成功", ctx)
}
