package pkg

import (
	"go.uber.org/zap"
	"nebula.xyz/model/system"
	"strings"

	"gorm.io/gorm"
	"nebula.xyz/global"
	"nebula.xyz/pkg/internal"
)

func Gorm() *gorm.DB {
	global.Logger.Error("初始化数据库开始", zap.String("配置文件数据库类型为：", global.CONFIG.SERVER.DbType))
	switch strings.ToLower(global.CONFIG.SERVER.DbType) {
	case "mysql":
		global.Logger.Info("系统当前数据库为MySQL")
		return internal.GormMySQL()
	case "pgsql":
		global.Logger.Info("系统当前数据库为Postgresql")
		return internal.GormPgsql()
	case "sqlite":
		global.Logger.Info("系统当前数据库为sqlite")
		return internal.GormSqlite()
	default:
		global.Logger.Info("系统当前数据库为sqlite")
		return internal.GormSqlite()
	}
}

func RegisterTables() {
	db := global.DB
	global.Logger.Info("初始化数据库表中...")
	db.AutoMigrate(
		system.Device{},
		system.DeviceChannel{},
		system.SipServer{},
		system.User{},
	)
}
