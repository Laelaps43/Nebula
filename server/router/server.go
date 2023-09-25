package router

import (
	"fmt"

	"go.uber.org/zap"
	"nebula.xyz/global"
)

// 运行整个系统
func RunServer() {

	Router := Routers()

	port := fmt.Sprintf(":%d", global.CONFIG.SERVER.PORT) // gin启动的地址和端口号

	Router.Run(port)
	global.Logger.Info("初始化gin完毕，系统运行在", zap.String("端口", port))
}
