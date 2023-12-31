package response

type Permission struct {
	Auths   []string `json:"auths"`
	Modules []string `json:"modules"`
}
