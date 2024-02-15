package helper

// 系统常量

const (
	ServerName = "Nebula"
	// 配置文件名称
	ConfigName = "nebula"
	// 配置文件类型
	ConfigType = "yaml"
	// 配置文件路径，当前工作路径
	ConfigPath = "./"

	// 日志文件名
	LogName = "nebula.log"

	// 系统使用Redis
	CacheRedis = 0

	// 系统使用其他
	CacheOther = 1

	CacheServerUpTimeKey = "serverUpTime"

	KeepTTL = -1

	// Sip服务器状态

	SipServerON  = 1
	SipServerOFF = 0

	// 流状态
	StreamClose = 1
	StreamStart = 2

	StreamUnRecord = 1 // 流未被录制

	StreamRecorded = 2 // 流正在被录制

	StreamApp = "rtp" // 国标固定为rtp

	// 录制类型
	RecordHLS = 0

	RecordMP4 = 1

	RecordMaxSecond = 0

	RecordPath = "/opt/media/video"

	// 设备状态
	DeviceOnline  = "1"
	DeviceOffline = "0"

	ChannelStatusON  = "ON"
	ChannelStatusOFF = "OFF"

	MediaStatusON = 1

	ZLMeidaHookSuccess = 0
	ZLMeidaHookFail    = 1

	ZLMeidaHookFailMessage    = "parser on_play_hook interface param fail, auth fail and not allow play"
	ZLMeidaHookSuccessMessage = "success"

	// 系统设备默认是传输协议
	DefaultTransPort = "UDP"

	// 系统默认最大删除次数
	MaxDeleteCount = 10
)

const (
	ZlmGetApiList    = "getApiList"
	ZlmStartRecord   = "startRecord"
	ZlmIsRecording   = "isRecording"
	ZlmStopRecording = "stopRecord"
	ZlmLoadMP4File   = "loadMP4File"
)

type QueryType string

const (
	DeviceStatusCmdType   QueryType = "DeviceStatus"
	CatalogCmdType        QueryType = "Catalog"
	DeviceInfoCmdType     QueryType = "DeviceInfo"
	RecordInfoCmdType     QueryType = "RecordInfo"
	AlarmCmdType          QueryType = "Alarm"
	ConfigDownloadCmdType QueryType = "ConfigDownload"
	PresetQueryCmdType    QueryType = "PresetQuery"
	MobilePositionCmdType QueryType = "MobilePosition"
	KeepaliveCmdType      QueryType = "Keepalive"
)
