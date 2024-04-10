package response

// 登录返回状态

type LoginResponse struct {
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expiresAt"`
}

type UserRole struct {
	Label string `json:"label"`
	Value uint   `json:"value"`
	Hold  bool   `json:"hold"`
}
