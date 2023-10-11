package request

type ChannelGenerate struct {
	DeviceId   string `json:"device_id"`
	ChannelNum string `json:"channel_num"`
}
