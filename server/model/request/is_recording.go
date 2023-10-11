package request

type IsRecording struct {
	Secret string `json:"secret"` // api操作密钥(配置文件配置)
	Type   string `json:"type"`   // 0为hls，1为mp4
	Vhost  string `json:"vhost"`  // 虚拟主机，例如__defaultVhost__
	App    string `json:"app"`    // 应用名，例如 live
	Stream string `json:"stream"` // 流id，例如 obs
}
