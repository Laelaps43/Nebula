package model

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 响应消息体

type Response struct{
	Code 		int 		`json:"code"`				// Http 状态码
	Data		any			`json:"data"`				// 返回的数据
	Message 	string		`json:"message"`			// 返回消息
}

const (
	ERROR = 0 			// 失败
	SUCCESS = 1 		// 成功
)

// 需要网前端返回数据
func Result(code int, data any, message string, c *gin.Context){
	c.JSON(http.StatusOK, Response{
		code,
		data,
		message,
	})
}

// 操作成功
func OK(c *gin.Context){
	Result(SUCCESS, nil, "操作成功", c)
}

// 操作成功，并自定义返回消息
func OKWithMessage(message string,c  *gin.Context){
	Result(SUCCESS, nil, message , c)
}

// 操作失败同时返回错误信息
func ErrorWithMessage(message string,  c *gin.Context){
	Result(ERROR, nil, message, c)
}

// 操作失败
func Error(c *gin.Context){
	Result(ERROR, nil, "操作失败", nil)
}