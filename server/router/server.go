package router

import (
	"fmt"
	"go.uber.org/zap"
	"nebula.xyz/global"
	"nebula.xyz/helper"
	"time"
)

// RunServer 运行整个系统
func RunServer() {

	Router := Routers()

	port := fmt.Sprintf(":%d", global.CONFIG.SERVER.PORT) // gin启动的地址和端口号

	// 保存系统启动时间
	_, err := global.CACHE.Set(helper.CacheServerUpTimeKey, time.Now().UnixMilli(), helper.KeepTTL)
	if err != nil {
		global.Logger.Error("保存系统启动时间失败", zap.Error(err))
	}
	err = Router.Run(port)
	if err != nil {
		global.Logger.Error("启动服务失败", zap.Error(err))
	}
}
