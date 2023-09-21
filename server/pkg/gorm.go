package pkg

import (
	"strings"

	"gorm.io/gorm"
	"nebula.xyz/global"
	"nebula.xyz/pkg/internal"
)



func Gorm() *gorm.DB{
	switch strings.ToLower(global.CONFING.SERVER.DbType) {
	case "mysql":
		return internal.GormMySQL();
	case "pgsql":
		return internal.GormPgsql()
	case "sqlite":
		return internal.GormSqlite()
	default:
		return internal.GormSqlite()
	}
}
