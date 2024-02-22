package models

import (
	"time"

	categoryModel "InitiaRe-website/internal/category/models"
	commonModel "InitiaRe-website/internal/models"
	"InitiaRe-website/pkg/utils"

	"github.com/jinzhu/copier"
)

type RequestList struct {
	commonModel.RequestPaging
	Email       string `json:"email"`
	Title       string `json:"title"`
	CategoryId  int    `json:"category_id"`
	CategoryIds string `json:"category_ids"`
	TypeId      int    `json:"type_id"`
	StatusId    int    `json:"status_id"`
	FromDate    int    `json:"from_date"`
	ToDate      int    `json:"to_date"`
	CreatedBy   int    `json:"created_by"`
}

func (r *RequestList) ToMap() map[string]interface{} {

	categoryIds := []int{}
	if r.CategoryIds != "" {
		categoryIds = utils.StringToArrayInt(r.CategoryIds, ",")
	}
	if r.CategoryId != 0 {
		categoryIds = append(categoryIds, r.CategoryId)
	}

	return map[string]interface{}{
		"from_date":    r.FromDate,
		"to_date":      r.ToDate,
		"created_by":   r.CreatedBy,
		"status_id":    r.StatusId,
		"type_id":      r.TypeId,
		"title":        r.Title,
		"email":        r.Email,
		"category_ids": categoryIds,
		"page":         r.Page,
		"size":         r.Size,
		"sort_by":      r.SortBy,
		"order_by":     r.OrderBy,
	}
}

type Response struct {
	Id                int    `json:"id"`
	CategoryId        int    `json:"category_id"`
	StatusId          int    `json:"status_id"`
	Title             string `json:"title"`
	ShortBrief        string `json:"short_brief,omitempty"`
	Content           string `json:"content,omitempty"`
	Thumbnail         string `json:"thumbnail,omitempty"`
	PrePublishContent string `json:"pre_publish_content,omitempty"`
	PublishDate       string `json:"publish_date,omitempty"`
	TypeId            int    `json:"type_id,omitempty"`
	TypeName          string `json:"type_name,omitempty"`
	CreatedAt         string `json:"created_at"`
	CreatedBy         int    `json:"created_by"`
	UpdatedAt         string `json:"updated_at,omitempty"`
	UpdatedBy         int    `json:"updated_by,omitempty"`

	// Custom fields
	StatusName    string                    `json:"status_name,omitempty"`
	CategoryName  string                    `json:"category_name,omitempty"`
	Email         string                    `json:"email,omitempty"`
	SubCategories []*categoryModel.Response `json:"sub_categories,omitempty"`
	Rating        int                       `json:"rating,omitempty"`
}

type SaveRequest struct {
	Id                int       `json:"id"`
	CategoryId        int       `json:"category_id"`
	SubCategoryIds    []int     `json:"sub_category_ids"`
	Content           string    `json:"content"`
	Title             string    `json:"title"`
	ShortBrief        string    `json:"short_brief"`
	Thumbnail         string    `json:"thumbnail"`
	PrePublishContent string    `json:"pre_publish_content"`
	TypeId            int       `json:"type_id"`
	PublishDate       time.Time `json:"publish_date"`
}

type ListPaging struct {
	commonModel.ListPaging
	Records []*Response
}

type ApprovedList struct {
	Total   int
	Records []*Response
}

type CreateRequest struct {
	Content           string `json:"content"`
	CategoryId        int    `json:"category_id"`
	SubCategoryIds    string `json:"sub_category_ids"`
	Title             string `json:"title"`
	ShortBrief        string `json:"short_brief"`
	Thumbnail         string `json:"thumbnail"`
	PrePublishContent string `json:"pre_publish_content"`
	TypeId            int    `json:"type_id"`
	PublishDate       string `json:"publish_date"`
}

func (r *CreateRequest) ToSaveRequest() *SaveRequest {
	req := &SaveRequest{}
	copier.Copy(req, r)

	if r.SubCategoryIds != "" {
		req.SubCategoryIds = utils.StringToArrayInt(r.SubCategoryIds, ",")
	}
	if r.PublishDate != "" {
		req.PublishDate, _ = time.Parse("2006-01-02 15:04:05", r.PublishDate)
	}

	return req
}

type UpdateRequest struct {
	Content           string    `json:"content"`
	CategoryId        int       `json:"category_id"`
	Title             string    `json:"title"`
	ShortBrief        string    `json:"short_brief"`
	Thumbnail         string    `json:"thumbnail"`
	PrePublishContent string    `json:"pre_publish_content"`
	TypeId            int       `json:"type_id"`
	PublishDate       time.Time `json:"publish_date"`
}

type ApproveRequest struct {
	Id int `json:"id"`
}

type DisableRequest struct {
	Id int `json:"id"`
}

func (r *UpdateRequest) ToSaveRequest(id int) *SaveRequest {
	req := &SaveRequest{}
	copier.Copy(req, r)
	req.Id = id
	return req
}
