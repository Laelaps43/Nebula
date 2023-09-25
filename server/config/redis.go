package config

// Redis 配置类

type Redis struct {
	Addr     string `mapstructure:"addr" yaml:"addr"`
	Password string `mapstructure:"password" yaml:"password" default:""`
	DB       int    `mapstructure:"db" yaml:"db" default:"0"`
}
