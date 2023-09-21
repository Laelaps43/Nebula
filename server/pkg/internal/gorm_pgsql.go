package internal

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"nebula.xyz/global"
)



func GormPgsql() *gorm.DB {
	// 获取Postgresql配置
	p := global.CONFING.PGSQL
	if p.DBName == "" {
		return nil
	}
	config := postgres.Config{
		DSN: p.Dsn(),
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}
	if db,err := gorm.Open(postgres.New(config)); err != nil{
		return nil
	}else{
		pgsql,_ := db.DB()
		pgsql.SetMaxIdleConns(p.MaxIdleConns)
		pgsql.SetMaxOpenConns(p.MaxOpenConns)
		return db
	}
}