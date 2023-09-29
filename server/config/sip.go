package config

// sip 配置文件，如果数据库中存在，则会从数据库中导入

type Sip struct {
	Ip            string `mapstructure:"ip" yaml:"ip"`                         // SIP 服务地址
	Port          uint   `mapstructure:"port" yaml:"port"`                     // SIP 端口
	Region        string `mapstructure:"region" yaml:"region"`                 // SIP 区域
	SipId         string `mapstructure:"sip-id" yaml:"sip-id"`                 // SIP 服务器ID
	Password      string `mapstructure:"password" yaml:"password"`             // SIP 密码
	UserAgent     string `mapstructure:"user-agent" yaml:"user-agent"`         // 代理
	DevicePrefix  string `mapstructure:"device-prefix" yaml:"device-prefix"`   // 设备前缀，用来添加设备
	ChannelPrefix string `mapstructure:"channel-prefix" yaml:"channel-prefix"` // 通道前缀，同来后续添加通道
}
