package internal

import (
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"nebula.xyz/global"
)

func GormMySQL() *gorm.DB {
	global.Logger.Info("系统连接MySQL中")
	m := global.CONFIG.MySQL
	if m.DBName == "" {
		global.Logger.Error("连接的数据库为", zap.String("DBName", m.DBName))
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(),
		DefaultStringSize:         256,
		SkipInitializeWithVersion: false,
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig)); err != nil {
		return nil
	} else {
		sqldb, _ := db.DB()
		sqldb.SetMaxOpenConns(m.MaxOpenConns)
		sqldb.SetMaxIdleConns(m.MaxIdleConns)
		global.Logger.Info("连接数据库成功")
		return db
	}
}
