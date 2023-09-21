package config
// 数据库基本信息

type DsnProvider interface{
	Dsn() string
}

type GeneralDB struct{
	Port			string		`yaml:"port"`				// 端口
	UserName		string		`yaml:"username"`			// 数据库用户名
	Password		string		`yaml:"password"`			// 密码
	DBName			string		`yaml:"db-name"`			// 数据库名
	Path			string		`yaml:"path"`				// 数据库url
	MaxIdleConns	int			`yaml:"max-idle-conns"`		// 空闲最大连接数
	MaxOpenConns	int 		`yaml:"max-open-conns"`		// 数据库最大连接数
	LogModel		string		`yaml:"log-model"`			// 开启Gorm全局日志等级 "silent"、"error"、"warn"、"info" 不填默认info 填入silent可以关闭控制台日志
	LogZap			bool		`yaml:"log-zap"`			// 是否将日志写入zap中
}