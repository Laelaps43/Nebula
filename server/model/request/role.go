package request

// CreateRole 创建角色
type CreateRole struct {
	Name     string `json:"name"`
	Slug     string `json:"slug"`
	Desc     string `json:"desc"`
	ParentId *uint  `json:"parentId"`
}

type UpdateRole struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}

type UpdatePermission struct {
	Menu   []uint `json:"menu"`
	Button []uint `json:"button"`
	RoleId uint64 `json:"roleId"`
}
