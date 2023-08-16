package models

import (
	commonModel "github.com/Ho-Minh/InitiaRe-website/internal/models"
	"github.com/jinzhu/copier"
)

type RequestList struct {
	commonModel.RequestPaging
	FromDate  int `json:"from_date"`
	ToDate    int `json:"to_date"`
	CreatedBy int `json:"created_by"`
}

func (r *RequestList) ToMap() map[string]interface{} {
	return map[string]interface{}{
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
	Id        int    `json:"id"`
	Category  string `json:"category"`
	CreatedAt string `json:"created_at"`
	CreatedBy int    `json:"created_by"`
	UpdatedAt string `json:"updated_at"`
	UpdatedBy int    `json:"updated_by"`
}

type SaveRequest struct {
	Id       int    `json:"id"`
	Category string `json:"category"`
}

type ListPaging struct {
	commonModel.ListPaging
	Records []*Response
}

type CreateRequest struct {
	Category string `json:"category"`
}

type UpdateRequest struct {
	Category string `json:"category"`
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