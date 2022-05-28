package structs

type Pagination struct {
	TotalRows int64       `json:"total_rows"`
	Data      interface{} `json:"data"`
}
