package response

import "nebula.xyz/model/system"

type GenerateInfo struct {
	Device     *system.Device          `json:"device"`
	Channels   []*system.DeviceChannel `json:"channels"`
	ChannelSum int                     `json:"channel_sum"`
}
