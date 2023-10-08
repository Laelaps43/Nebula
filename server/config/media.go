package config

type Media struct {
	RTP      string `mapstructure:"rtp" yaml:"rtp"`
	Restful  string `mapstructure:"restful" yaml:"restful"`
	Secret   string `mapstructure:"secret" yaml:"secret"`
	Address  string `mapstructure:"address" yaml:"address"`
	RTSPPort string `mapstructure:"rtsp-port" yaml:"rtsp-port"`
	RTMPPort string `mapstructure:"rtmp-port" yaml:"rtmp-port"`
}
