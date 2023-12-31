package zlm

type OnPublish struct {
	App           string `json:"app"`
	HookIndex     int    `json:"hook_index"`
	Id            string `json:"id"`
	IP            string `json:"ip"`
	MediaServerId string `json:"mediaServerId"`
	OriginType    int    `json:"originType"`
	OriginTypeStr string `json:"originTypeStr"`
	Params        string `json:"params"`
	Port          int    `json:"port"`
	Schema        string `json:"schema"`
	Stream        string `json:"stream"`
	Vhost         string `json:"vhost"`
}
