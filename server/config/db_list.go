package config

// 数据库基本信息

type DsnProvider interface {
	Dsn() string
}

type GeneralDB struct {
	Port         string `mapstructure:"port" yaml:"port"`                   // 端口
	UserName     string `mapstructure:"username" yaml:"username"`           // 数据库用户名
	Password     string `mapstructure:"password" yaml:"password"`           // 密码
	DBName       string `mapstructure:"db-name" yaml:"db-name"`             // 数据库名
	Path         string `mapstructure:"path" yaml:"path"`                   // 数据库url
	MaxIdleConns int    `mapstructure:"max-idle-cons" yaml:"max-idle-cons"` // 空闲最大连接数
	MaxOpenConns int    `mapstructure:"max-open-cons" yaml:"max-open-cons"` // 数据库最大连接数
	LogModel     string `mapstructure:"log-model" yaml:"log-model"`         // 开启Gorm全局日志等级 "silent"、"error"、"warn"、"info" 不填默认info 填入silent可以关闭控制台日志
	LogZap       bool   `mapstructure:"log-zap" yaml:"log-zap"`             // 是否将日志写入zap中
}
