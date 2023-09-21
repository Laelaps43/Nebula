package config
// 系统服务配置

type SERVER struct{
	Mode 			string 		`yaml:"mode"`

	DbType			string		`yaml:"db_type`

	PORT			int 		`yaml:"port"`

	RouterPrefix	string	`yaml:"router-prefix"`
}