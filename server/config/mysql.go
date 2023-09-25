package config

// MySQL数据库数据

type MySQL struct {
	GeneralDB `yaml:",inline" mapstructure:",squash"`
}

func (m *MySQL) Dsn() string {
	return m.UserName + ":" + m.Password + "@tcp(" + m.Path + ":" + m.Port + ")/" + m.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
}
