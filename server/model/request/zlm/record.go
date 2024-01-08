package zlm

type StartRecord struct {
	Secret         string `json:"secret"`
	Type           int    `json:"type"`
	Vhost          string `json:"vhost"`
	App            string `json:"app"`
	Stream         string `json:"stream"`
	CustomizedPath string `json:"customized_path,omitempty"`
	MaxSecond      int    `json:"max_second,omitempty"`
}

type StopRecord struct {
	Secret string `json:"secret"` // api操作密钥(配置文件配置)
	Type   int    `json:"type"`   // 0为hls，1为mp4
	Vhost  string `json:"vhost"`  // 虚拟主机，例如__defaultVhost__
	App    string `json:"app"`    // 应用名，例如 live
	Stream string `json:"stream"` // 流id，例如 obs
}

type RecordMp4 struct {
	MediaServerId string  `json:"mediaServerId"`
	App           string  `json:"app"`
	FileName      string  `json:"file_name"`
	FilePath      string  `json:"file_path"`
	FileSize      int64   `json:"file_size"`
	Folder        string  `json:"folder"`
	StartTime     int64   `json:"start_time"`
	Stream        string  `json:"stream"`
	TimeLen       float64 `json:"time_len"`
	URL           string  `json:"url"`
	Vhost         string  `json:"vhost"`
}

type IsRecording struct {
	Secret string `json:"secret"` // api操作密钥(配置文件配置)
	Type   int    `json:"type"`   // 0为hls，1为mp4
	Vhost  string `json:"vhost"`  // 虚拟主机，例如__defaultVhost__
	App    string `json:"app"`    // 应用名，例如 live
	Stream string `json:"stream"` // 流id，例如 obs
}

type PlayRecord struct {
	Secret     string `json:"secret"`      // api操作密钥(配置文件配置)
	Vhost      string `json:"vhost"`       // 虚拟主机，例如__defaultVhost__
	App        string `json:"app"`         // 应用名，例如 live
	Stream     string `json:"stream"`      // 流id，例如 obs
	FilePath   string `json:"file_path"`   // 文件路径
	FileRepeat int    `json:"file_repeat"` // 是否循环播放 0-否 1-是
}
