package main

import (
	"go.uber.org/zap"
	"nebula.xyz/global"
	"nebula.xyz/pkg"
	"nebula.xyz/router"
	"nebula.xyz/sip"
	"nebula.xyz/utils"
)

func main() {
	global.VP = pkg.Viper()
	global.Logger = pkg.Zap()
	zap.ReplaceGlobals(global.Logger)
	global.DB = pkg.Gorm()
	global.Logger.Error("数据库的值为", zap.Bool("db", global.DB != nil))
	if global.DB != nil {
		// TODO 初始化表
		pkg.RegisterTables()
	}
	global.CACHE = pkg.GetCache()
	sip.InitSipServer()
	global.Logger.Info("获取系统信息.....")
	go utils.GetSystemInfo()
	go utils.GetSystemInfo()
	router.RunServer()
}
