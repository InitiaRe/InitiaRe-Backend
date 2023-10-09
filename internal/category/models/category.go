package models

import (
	commonModel "github.com/Ho-Minh/InitiaRe-website/internal/models"
	"github.com/jinzhu/copier"
)

type RequestList struct {
	commonModel.RequestPaging
	ArticleId int	`json:"article_id"`
	FromDate  int	`json:"from_date"`
	ToDate    int	`json:"to_date"`
	CreatedBy int	`json:"created_by"`
}

func (r *RequestList) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"article_id": r.ArticleId,
		"from_date":  r.FromDate,
		"to_date":    r.ToDate,
		"created_by": r.CreatedBy,
		"page":       r.Page,
		"size":       r.Size,
		"sort_by":    r.SortBy,
		"order_by":   r.OrderBy,
	}
}

type Response struct {
	Id           int    `json:"id"`
	CategoryName string `json:"category_name"`
	CreatedAt    string `json:"created_at"`
	CreatedBy    int    `json:"created_by"`
	UpdatedAt    string `json:"updated_at,omitempty"`
	UpdatedBy    int    `json:"updated_by,omitempty"`
}

type SaveRequest struct {
	Id           int    `json:"id"`
	CategoryName string `json:"category_name"`
}

type ListPaging struct {
	commonModel.ListPaging
	Records []*Response
}

type CreateRequest struct {
	CategoryName string `json:"category_name"`
}

type UpdateRequest struct {
	CategoryName string `json:"category_name"`
}

func (r *CreateRequest) ToSaveRequest() *SaveRequest {
	req := &SaveRequest{}
	copier.Copy(req, r)
	return req
}

func (r *UpdateRequest) ToSaveRequest(id int) *SaveRequest {
	req := &SaveRequest{}
	copier.Copy(req, r)
	req.Id = id
	return req
}
