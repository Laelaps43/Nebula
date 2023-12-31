package model

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 响应消息体
type Response struct {
	Code    int    `json:"code"`    // Http 状态码 1正确返回， 0 为错误返回
	Data    any    `json:"data"`    // 返回的数据
	Message string `json:"message"` // 返回消息
}

const (
	ERROR   = 0 // 失败
	SUCCESS = 1 // 成功
)

// 需要网前端返回数据
func result(code int, data any, message string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		message,
	})
}

// OkWithDetailed 存在成功，携带数据
func OkWithDetailed(object any, message string, c *gin.Context) {
	result(SUCCESS, object, message, c)
}

// OK 操作成功
func OK(c *gin.Context) {
	result(SUCCESS, nil, "操作成功", c)
}

// OKWithMessage 操作成功，并自定义返回消息
func OKWithMessage(message string, c *gin.Context) {
	result(SUCCESS, nil, message, c)
}
func ErrorWithDetailed(object any, message string, c *gin.Context) {
	result(ERROR, object, message, c)
}

// ErrorWithMessage 操作失败同时返回错误信息
func ErrorWithMessage(message string, c *gin.Context) {
	result(ERROR, nil, message, c)
}

// Error 操作失败
func Error(c *gin.Context) {
	result(ERROR, nil, "操作失败", nil)
}

// ServerError 服务器内部错误500
func ServerError(c *gin.Context) {
	// 返回500错误
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": "Internal Server Error",
	})
}
