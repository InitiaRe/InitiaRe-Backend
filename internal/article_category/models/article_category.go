package models

import (
	commonModel "github.com/Ho-Minh/InitiaRe-website/internal/models"
)

type RequestList struct {
	commonModel.RequestPaging
	ArticleId  int
	CategoryId int
	Type       int
}

func (r *RequestList) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"article_id":  r.ArticleId,
		"category_id": r.CategoryId,
		"type":        r.Type,
		"page":        r.Page,
		"size":        r.Size,
		"sort_by":     r.SortBy,
		"order_by":    r.OrderBy,
	}
}

type Response struct {
	Id         int `json:"id"`
	CategoryId int `json:"category_id"`
	ArticleId  int `json:"article_id"`
	Type       int `json:"type"`
}

type SaveRequest struct {
	Id         int `json:"id"`
	CategoryId int `json:"category_id"`
	ArticleId  int `json:"article_id"`
	Type       int `json:"type"`
}

type ListPaging struct {
	commonModel.ListPaging
	Records []*Response
}
