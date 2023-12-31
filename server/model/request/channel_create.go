package request

type CreateChannel struct {
	Name      string `json:"name"`
	ChannelId string `json:"channelId"`
	DeviceId  string `json:"deviceId"`
}
