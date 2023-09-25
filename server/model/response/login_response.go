package response

// 登录返回状态

import (
	"nebula.xyz/model/system"
)

type LoginResponse struct {
	User      *system.User `json:"user"`
	Token     string       `json:"token"`
	ExpiresAt int64        `json:"expiresAt"`
}
