package response

type RecordPageResponse struct {
	Stream         string `json:"stream"`
	ChannelName    string `json:"channelName"`
	ChannelID      string `json:"channelId"`
	DeviceName     string `json:"deviceName"`
	IsRecording    uint   `json:"isRecording"`
	LastRecordTime string `json:"lastRecordTime"`
	Duration       int64  `json:"duration"`
}

type RecordSelect struct {
	Label string `json:"label"`
	Value uint   `json:"value"`
}
