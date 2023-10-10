package request

type ServerKeepalive struct {
	MediaServerId string              `json:"mediaServerId,omitempty"`
	Data          ServerKeepaliveData `json:"data"`
}

type ServerKeepaliveData struct {
	Buffer                int `json:"Buffer,omitempty"`
	BufferLikeString      int `json:"BufferLikeString,omitempty"`
	BufferList            int `json:"BufferList,omitempty"`
	BufferRaw             int `json:"BufferRaw,omitempty"`
	Frame                 int `json:"Frame,omitempty"`
	FrameImp              int `json:"FrameImp,omitempty"`
	MediaSource           int `json:"MediaSource,omitempty"`
	MultiMediaSourceMuxer int `json:"MultiMediaSourceMuxer,omitempty"`
	RtmpPacket            int `json:"RtmpPacket,omitempty"`
	RtpPacket             int `json:"RtpPacket,omitempty"`
	Socket                int `json:"Socket,omitempty"`
	TcpClient             int `json:"TcpClient,omitempty"`
	TcpServer             int `json:"TcpServer,omitempty"`
	TcpSession            int `json:"TcpSession,omitempty"`
	UdpServer             int `json:"UdpServer,omitempty"`
	UdpSession            int `json:"UdpSession,omitempty"`
}
