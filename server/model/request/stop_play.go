package request

type StopPlay struct {
	DeviceId  string `json:"device_id"`
	ChannelId string `json:"channel_id"`
}
