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

	// Sip服务器状态
	SipServerON  = 1
	SipServerOFF = 0

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
)

const (
	ZlmGetApiList  = "getApiList"
	ZlmStartRecord = "startRecord"
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
