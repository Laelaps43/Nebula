package response

type OverViewResult struct {
	OnlineDevice  int64 `json:"onlineDevice"`
	OfflineDevice int64 `json:"offlineDevice"`
	Channel       int64 `json:"channel"`
	Video         int64 `json:"video"`
}

type ServerInfo struct {
	ServerDetails      `json:"serverDetails"`
	MediaServerDetails `json:"mediaServerDetails"`
}

type ServerDetails struct {
	ServiceAddress  string `json:"serviceAddress"`
	SIPServerID     string `json:"sipServerID"`
	SIPServerDomain string `json:"sipServerDomain"`
	SIPPassword     string `json:"sipPassword"`
	Uptime          int64  `json:"uptime"`
}

type MediaServerDetails struct {
	MediaServiceAddress string `json:"mediaServiceAddress"`
	MediaUniqueID       string `json:"mediaUniqueID"`
	RTPPort             string `json:"rtpPort"`
	RestfulPort         string `json:"restfulPort"`
	RTSPPort            string `json:"rtspPort"`
	RTMPPort            string `json:"rtmpPort"`
	TCPSessions         int    `json:"tcpSessions"`
	UDPSessions         int    `json:"udpSessions"`
	LastHeartbeatTime   string `json:"lastHeartbeatTime"`
}
