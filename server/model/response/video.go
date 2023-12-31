package response

// PlayResponsePayload 点播返回
type PlayResponsePayload struct {
	HTTP  string `gorm:"comment:mu38地址"`
	RTMP  string `gorm:"comment:RTMP地址"`
	RTSP  string `gorm:"comment:RTSP地址"`
	WSFLV string `gorm:"comment:WSFLV地址"`
}
