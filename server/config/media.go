package config

type Media struct {
	RTP              string  `mapstructure:"rtp" yaml:"rtp"`
	Restful          string  `mapstructure:"restful" yaml:"restful"`
	Secret           string  `mapstructure:"secret" yaml:"secret"`
	Address          string  `mapstructure:"address" yaml:"address"`
	RTSPPort         string  `mapstructure:"rtsp-port" yaml:"rtsp-port"`
	RTMPPort         string  `mapstructure:"rtmp-port" yaml:"rtmp-port"`
	MediaServerId    string  `mapstructure:"media-server-id" yaml:"media-server-id"`
	RecordPath       string  `mapstructure:"record-path" yaml:"record-path"`
	StorageThreshold float64 `mapstructure:"storage-threshold" yaml:"storage-threshold"`
	Domain           string  `mapstructure:domain yaml:"domain"`
}
