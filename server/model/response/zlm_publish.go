package response

type OnPublish struct {
	Code           int    `json:"code" default:"0"`
	Msg            string `json:"msg" default:""`
	EnableHLS      bool   `json:"enable_hls" default:"true"`
	EnableHLSFMP4  bool   `json:"enable_hls_fmp4" default:"true"`
	EnableMP4      bool   `json:"enable_mp4" default:"true"`
	EnableRTSP     bool   `json:"enable_rtsp" default:"true"`
	EnableRTMP     bool   `json:"enable_rtmp" default:"true"`
	EnableTS       bool   `json:"enable_ts" default:"true"`
	EnableFMP4     bool   `json:"enable_fmp4" default:"true"`
	HLSDemand      bool   `json:"hls_demand" default:"false"`
	RTSPDemand     bool   `json:"rtsp_demand" default:"false"`
	RTMPDemand     bool   `json:"rtmp_demand" default:"false"`
	TSDemand       bool   `json:"ts_demand" default:"false"`
	FMP4Demand     bool   `json:"fmp4_demand" default:"false"`
	EnableAudio    bool   `json:"enable_audio" default:"false"`
	AddMuteAudio   bool   `json:"add_mute_audio" default:"false"`
	MP4SavePath    string `json:"mp4_save_path" default:"./mp4_save_path/"`
	MP4MaxSecond   int    `json:"mp4_max_second" default:"3600"`
	MP4AsPlayer    bool   `json:"mp4_as_player" default:"false"`
	HLSSavePath    string `json:"hls_save_path" default:"./hls_save_path/"`
	ModifyStamp    int    `json:"modify_stamp" default:"0"`
	ContinuePushMS uint32 `json:"continue_push_ms" default:"30000"`
	AutoClose      bool   `json:"auto_close" default:"false"`
	StreamReplace  string `json:"stream_replace" default:""`
}
