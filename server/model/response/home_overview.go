package response

type OverViewResult struct {
	OnlineDevice  int64 `json:"onlineDevice"`
	OfflineDevice int64 `json:"offlineDevice"`
	Channel       int64 `json:"channel"`
	Video         int64 `json:"video"`
}
