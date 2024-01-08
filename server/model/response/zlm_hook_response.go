package response

type ZLMHookResponse struct {
	// 0代表允许，其他均为不允许
	Code int `json:"code"`

	// 当code不为0时，msg字段应给出相应提示
	Msg string `json:"msg"`
}
