package models

import (
	commonModel "InitiaRe-website/internal/models"

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
	Content   string `json:"content,omitempty"`
	CreatedBy int    `json:"created_by,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedBy int    `json:"updated_by,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

type SaveRequest struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
}

type CreateRequest struct {
	Content string `json:"content"`
}

func (r *CreateRequest) ToSaveRequest() *SaveRequest {
	req := &SaveRequest{}
	copier.Copy(req, r)
	return req
}

type UpdateRequest struct {
	Content string `json:"content"`
}

func (r *UpdateRequest) ToSaveRequest(id int) *SaveRequest {
	req := &SaveRequest{}
	copier.Copy(req, r)
	req.Id = id
	return req
}

type ListPaging struct {
	commonModel.ListPaging
	Records []*Response
}
