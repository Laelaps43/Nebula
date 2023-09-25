package config

// 系统JWT配置

type JWT struct {
	SigningKey string `mapstructure:"signing-key" yaml:"signing-key""`
	JwtExpire  string `mapstructure:"jwt-expire" yaml:"jwt-expire"`
	Issuer     string `mapstructure:"issuer" yaml:"issuer"`
}
