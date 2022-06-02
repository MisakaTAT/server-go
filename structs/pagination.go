package structs

// Pagination 分页封装
type Pagination struct {
	TotalRows int64       `json:"total_rows"`
	Data      interface{} `json:"data"`
}
