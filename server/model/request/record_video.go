package request

type RecordVideo struct {
	ChannelId string `json:"channelId"`
	DeviceId  string `json:"deviceId"`
}

type RecordRange struct {
	Start  string `json:"start"`
	End    string `json:"end"`
	Stream string `json:"stream"`
}
