package config

// 系统服务配置

type SERVER struct {
	Mode string `mapstructure:"mode" yaml:"mode"`

	DbType string `mapstructure:"db-type" yaml:"db-type"`

	PORT int `mapstructure:"port" yaml:"port"`

	RouterPrefix string `mapstructure:"router-prefix" yaml:"router-prefix"`

	LoginMaxNum uint `mapstructure:"login-max-num" yaml:"login-max-num"`

	LoginTimeout uint `mapstructure:"login-timeout" yaml:"login-timeout"`

	CacheType string `mapstructure:"cache-type" yaml:"cache-type"`
}
