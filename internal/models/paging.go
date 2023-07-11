package models

type RequestPaging struct {
	Page    int    `json:"page"`
	Size    int    `json:"size"`
	SortBy  string `json:"sort_by"`
	OrderBy string `json:"order_by"`
}

type ListPaging struct {
	Page    int
	Size    int
	Total   int
	Records interface{}
}
