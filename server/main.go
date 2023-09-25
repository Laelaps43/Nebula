package main

import (
	"go.uber.org/zap"
	"nebula.xyz/global"
	"nebula.xyz/pkg"
	"nebula.xyz/router"
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
	router.RunServer()
}
