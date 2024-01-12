package models

import (
	commonModel "github.com/Ho-Minh/InitiaRe-website/internal/models"
	"github.com/jinzhu/copier"
)

type RequestList struct {
	commonModel.RequestPaging
	ArticleId int `json:"article_id"`
	UserId    int `json:"user_id"`
	Rating    int `json:"rating"`
	Status    int `json:"status"`
}

func (r *RequestList) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"article_id": r.ArticleId,
		"user_id":    r.UserId,
		"rating":     r.Rating,
		"status":     r.Status,
		"page":       r.Page,
		"size":       r.Size,
		"sort_by":    r.SortBy,
		"order_by":   r.OrderBy,
	}
}

type Response struct {
	Id        int
	UserId    int
	ArticleId int
	Rating    int
	Status    int
}

type SaveRequest struct {
	Id        int
	UserId    int
	ArticleId int
	Rating    int
	Status    int
}

type ListPaging struct {
	commonModel.ListPaging
	Records []*Response
}

type VoteRequest struct {
	ArticleId int `json:"article_id"`
}

func (r *VoteRequest) ToSaveRequest(userId int) *SaveRequest {
	req := &SaveRequest{}
	copier.Copy(req, r)
	req.UserId = userId
	return req
}
