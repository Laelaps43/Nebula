package request

type StreamChangedData struct {
	Regist           bool   `json:"regist"`
	AliveSecond      int    `json:"aliveSecond"`
	App              string `json:"app"`
	BytesSpeed       int    `json:"bytesSpeed"`
	CreateStamp      int64  `json:"createStamp"`
	MediaServerID    string `json:"mediaServerId"`
	OriginTypeStr    string `json:"originTypeStr"`
	OriginURL        string `json:"originUrl"`
	ReaderCount      int    `json:"readerCount"`
	Schema           string `json:"schema"`
	Stream           string `json:"stream"`
	TotalReaderCount int    `json:"totalReaderCount"`
	VHost            string `json:"vhost"`
	OriginSock       struct {
		Identifier string `json:"identifier"`
		LocalIP    string `json:"local_ip"`
		LocalPort  int    `json:"local_port"`
		PeerIP     string `json:"peer_ip"`
		PeerPort   int    `json:"peer_port"`
	} `json:"originSock"`
	Tracks []struct {
		Channels   int    `json:"channels"`
		CoecID     int    `json:"codec_id"`
		CoecIDName string `json:"codec_id_name"`
		CoecType   int    `json:"codec_type"`
		Ready      bool   `json:"ready"`
		SampleBit  int    `json:"sample_bit"`
		SampleRate int    `json:"sample_rate"`
		FPS        int    `json:"fps"`
		Height     int    `json:"height"`
		Width      int    `json:"width"`
	} `json:"tracks"`
}
