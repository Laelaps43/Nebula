package request

// 登录字段
type Login struct {
	Email    string `json:"email"` // 用户名
	Password string `json:"password"`
}
