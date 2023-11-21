package models

import (
	commonModel "github.com/Ho-Minh/InitiaRe-website/internal/models"
)

type RequestList struct {
	commonModel.RequestPaging
	UserId int
}

func (r *RequestList) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"user_id":  r.UserId,
		"page":     r.Page,
		"size":     r.Size,
		"sort_by":  r.SortBy,
		"order_by": r.OrderBy,
	}
}

type Response struct {
	Id                 int `json:"id"`
	UserId             int `json:"user_id"`
	NumberUploaded     int `json:"number_uploaded"`
	NumberPeerReviewed int `json:"number_peer_reviewed"`
	NumberSpecReviewed int `json:"number_spec_reviewed"`
	Status             int `json:"status"`
	Role               int `json:"role"`
}

type SaveRequest struct {
	Id                 int
	UserId             int
	NumberUploaded     int
	NumberPeerReviewed int
	NumberSpecReviewed int
	Status             int
	Role               int
}

type ListPaging struct {
	commonModel.ListPaging
	Records []*Response
}
