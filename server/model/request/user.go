package request

// Login 登录字段
type Login struct {
	Email    string `json:"email"` // 用户名
	Password string `json:"password"`
}

// EnableUser 冻结用户
type EnableUser struct {
	ID     uint `json:"id"`
	Enable int  `json:"enable"`
}

type EditUser struct {
	ID       uint   `json:"id"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

type SwitchRole struct {
	RoleId uint `json:"roleId"`
}
