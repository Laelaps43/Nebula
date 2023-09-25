package web
// 登录路由

import "github.com/gin-gonic/gin"


type Login struct{}


// 登录处理
func (l *Login) LoginRouter(Router *gin.RouterGroup) (R gin.IRouter){
	LoginRouter := Router.Group("")
	return LoginRouter
}