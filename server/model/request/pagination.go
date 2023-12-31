package request

// Pagination 分页查询结构
type Pagination struct {
	Limit    int    `json:"limit"`
	Page     int    `json:"page"`
	DeviceId string `json:"deviceId"`
}
