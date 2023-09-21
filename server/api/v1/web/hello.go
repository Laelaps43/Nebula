package web

import (
	"github.com/gin-gonic/gin"
	"nebula.xyz/global"
	"nebula.xyz/model"
)



type Hello struct{}


func (h *Hello) SayHello(ctx *gin.Context){
	
	global.Logger.Info("执行SayHello函数")
	model.OKWithMessage("Hello!!!", ctx)
}