package internal

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"nebula.xyz/global"
)

func GormSqlite() *gorm.DB {
	q := global.CONFIG.SQLITE
	if q.DBName == "" {
		return nil
	}
	// 这里可以加入gorm配置
	if db, err := gorm.Open(sqlite.Open(q.Dsn()), nil); err != nil {
		return nil
	} else {
		sl, _ := db.DB()
		sl.SetMaxIdleConns(q.MaxIdleConns)
		sl.SetMaxOpenConns(q.MaxOpenConns)
		return db
	}
}
