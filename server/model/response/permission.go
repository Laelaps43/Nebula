package response

type Permission struct {
	Auths   []string `json:"auths"`
	Modules []string `json:"modules"`
}

// PermissionResponse 用户权限
type PermissionResponse struct {
	Menus   []PermissionDetails `json:"menus"`
	Buttons []PermissionDetails `json:"buttons"`
}

type PermissionDetails struct {
	Label   string `json:"label"`
	Value   uint   `json:"value"`
	Disable bool   `json:"disable"`
	Hold    bool   `json:"hold"`
}
