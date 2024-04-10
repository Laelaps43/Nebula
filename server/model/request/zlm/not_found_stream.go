package zlm

// StreamNotFound 表示一个流媒体应用的连接信息
type StreamNotFound struct {
	App           string `json:"app"`           // 流应用名
	ID            string `json:"id"`            // TCP链接唯一ID
	IP            string `json:"ip"`            // 播放器ip
	Params        string `json:"params"`        // 播放url参数
	Port          uint16 `json:"port"`          // 播放器端口号
	Schema        string `json:"schema"`        // 播放的协议，可能是rtsp、rtmp
	Stream        string `json:"stream"`        // 流ID
	Vhost         string `json:"vhost"`         // 流虚拟主机
	MediaServerID string `json:"mediaServerId"` // 服务器id,通过配置文件设置
}
