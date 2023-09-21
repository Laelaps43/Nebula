package main

import (
	"go.uber.org/zap"
	"nebula.xyz/global"
	"nebula.xyz/pkg"
)

func main()  {
	global.VP = pkg.Viper()	
	global.Logger = pkg.Zap()
	zap.ReplaceGlobals(global.Logger)
	global.DB = pkg.Gorm()
}