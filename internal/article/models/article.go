package models

import (
	"time"

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
	Id                int    `json:"id"`
	CategoryId        int    `json:"category_id"`
	StatusId          int    `json:"status_id"`
	Title             string `json:"title"`
	ShortBrief        string `json:"short_brief,omitempty"`
	Content           string `json:"content"`
	Thumbnail         string `json:"thumbnail,omitempty"`
	PrePublishContent string `json:"pre_publish_content,omitempty"`
	PublishDate       string `json:"publish_date"`
	CreatedAt         string `json:"created_at"`
	CreatedBy         int    `json:"created_by"`
	UpdatedAt         string `json:"updated_at"`
	UpdatedBy         int    `json:"updated_by"`
}

type SaveRequest struct {
	Id                int       `json:"id"`
	CategoryId        int       `json:"category_id"`
	Content           string    `json:"content"`
	Title             string    `json:"title"`
	ShortBrief        string    `json:"short_brief"`
	Thumbnail         string    `json:"thumbnail"`
	PrePublishContent string    `json:"pre_publish_content"`
	PublishDate       time.Time `json:"publish_date"`
}

type ListPaging struct {
	commonModel.ListPaging
	Records []*Response
}

type CreateRequest struct {
	Content           string    `json:"content"`
	CategoryId        int       `json:"category_id"`
	Title             string    `json:"title"`
	ShortBrief        string    `json:"short_brief"`
	Thumbnail         string    `json:"thumbnail"`
	PrePublishContent string    `json:"pre_publish_content"`
	PublishDate       time.Time `json:"publish_date"`
}

func (r *CreateRequest) ToSaveRequest() *SaveRequest {
	req := &SaveRequest{}
	copier.Copy(req, r)
	return req
}

type UpdateRequest struct {
	Content           string    `json:"content"`
	CategoryId        int       `json:"category_id"`
	Title             string    `json:"title"`
	ShortBrief        string    `json:"short_brief"`
	Thumbnail         string    `json:"thumbnail"`
	PrePublishContent string    `json:"pre_publish_content"`
	PublishDate       time.Time `json:"publish_date"`
}

func (r *UpdateRequest) ToSaveRequest(id int) *SaveRequest {
	req := &SaveRequest{}
	copier.Copy(req, r)
	req.Id = id
	return req
}
