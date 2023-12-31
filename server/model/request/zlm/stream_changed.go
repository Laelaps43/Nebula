package zlm

type originSock struct {
	Identifier string `json:"identifier"` //
	LocalIP    string `json:"local_ip"`   // 本机ip
	LocalPort  int    `json:"local_port"` // 本机端口
	PeerIP     string `json:"peer_ip"`    // 对端ip
	PeerPort   int    `json:"peer_port"`  // 对端port
}

type track struct {
	Channels    int    `json:"channels"`      // 音频通道数
	CoDecID     int    `json:"codec_id"`      // 264 = 0, H265 = 1, AAC = 2, G711A = 3, G711U = 4
	CoDecIDName string `json:"codec_id_name"` //  编码类型名称
	CoDecType   int    `json:"codec_type"`    // Video = 0, Audio = 1
	Ready       bool   `json:"ready"`         // 轨道是否准备就绪
	SampleBit   int    `json:"sample_bit"`    // 音频采样位数
	SampleRate  int    `json:"sample_rate"`   // 音频采样率
	FPS         int    `json:"fps"`           // 视频fps
	Height      int    `json:"height"`        // 视频高
	Width       int    `json:"width"`         // 视频宽
}

type StreamChange struct {
	Regist           bool       `json:"regist"`
	AliveSecond      int        `json:"aliveSecond"`      // 存活时间，单位秒
	App              string     `json:"app"`              // 应用名
	BytesSpeed       int        `json:"bytesSpeed"`       // 数据产生速度，单位byte/s
	CreateStamp      int64      `json:"createStamp"`      // GMT unix系统时间戳，单位秒
	MediaServerID    string     `json:"mediaServerId"`    // 服务器id
	OriginSock       originSock `json:"originSock"`       //
	OriginType       int        `json:"originType"`       // 产生源类型，包括 unknown = 0,rtmp_push=1,rtsp_push=2,rtp_push=3,pull=4,ffmpeg_pull=5,mp4_vod=6,device_chn=7,rtc_push=8
	OriginTypeStr    string     `json:"originTypeStr"`    //
	OriginURL        string     `json:"originUrl"`        // 产生源的url
	ReaderCount      int        `json:"readerCount"`      // 本协议观看人数
	Schema           string     `json:"schema"`           // 协议
	Stream           string     `json:"stream"`           // 流id
	TotalReaderCount int        `json:"totalReaderCount"` //  观看总人数，包括hls/rtsp/rtmp/http-flv/ws-flv/rtc
	Tracks           []track    `json:"tracks"`           //
	Vhost            string     `json:"vhost"`
}
