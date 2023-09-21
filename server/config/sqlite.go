package config
// sqlite配置

import "path/filepath"


type Sqlite struct{
	GeneralDB `yaml:",inline"`
}

// 获取sqlite的dsn
func (s Sqlite)Dsn() string{
	return filepath.Join(s.Path,s.DBName+".db")
}