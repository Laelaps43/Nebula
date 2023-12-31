package request

// CreateRole 创建角色
type CreateRole struct {
	Name     string `json:"name"`
	Slug     string `json:"slug"`
	Desc     string `json:"desc"`
	ParentId *uint  `json:"parentId"`
}
