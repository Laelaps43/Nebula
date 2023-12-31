package response

// PaginationResult 分页查询结果
type PaginationResult struct {
	List  any   `json:"list"`
	Total int64 `json:"total"`
}
