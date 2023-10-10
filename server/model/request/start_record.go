package request

type StartRecord struct {
	Secret         string `json:"secret"`          // api操作密钥(配置文件配置)
	Type           string `json:"type"`            // 0为hls，1为mp4
	Vhost          string `json:"vhost"`           // 虚拟主机，例如__defaultVhost__
	App            string `json:"app"`             // 应用名，例如 live
	Stream         string `json:"stream"`          // 流id，例如 obs
	CustomizedPath string `json:"customized_path"` // 录像保存目录
	MaxSecond      int    `json:"max_second"`      // mp4录像切片时间大小,单位秒，置0则采用配置项
}
