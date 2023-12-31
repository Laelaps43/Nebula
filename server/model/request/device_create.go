package request

type DeviceCreate struct {
	DeviceId string `json:"deviceId"`
	Name     string `json:"name"`
	Port     int    `json:"port"`
}
