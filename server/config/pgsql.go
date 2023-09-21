package config
// postgresql配置


type Pgsql struct{
	GeneralDB `yaml:",inline"`
}


func (p Pgsql)Dsn()string{
	return "user=" + p.UserName + " password=" + p.Password +
	" dbname=" + p.DBName + " port=" + p.Port + " sslmode=disable TimeZone=Asia/Shanghai"
}