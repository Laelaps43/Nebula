package internal

import (
	"gorm.io/gorm"
	"nebula.xyz/global"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
)


func GormMySQL() *gorm.DB{
	m := global.CONFING.MySQL
	if m.DBName == "" {
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN:	m.Dsn(),
		DefaultStringSize: 256,
		SkipInitializeWithVersion: false,
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig)); err != nil{
		return nil
	}else{
		sqldb, _ := db.DB()
		sqldb.SetMaxOpenConns(m.MaxOpenConns)
		sqldb.SetMaxIdleConns(m.MaxIdleConns)
		return db
	}
}