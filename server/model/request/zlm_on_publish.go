package request

type OnPublish struct {
	MediaServerId string `json:"mediaServerId"`
	App           string `json:"app"`
	Id            string `json:"id"`
	IP            string `json:"ip"`
	Params        string `json:"params"`
	Port          string `json:"port"`
	Schema        string `json:"schema"`
	Stream        string `json:"stream"`
	Vhost         string `json:"vhost"`
}
