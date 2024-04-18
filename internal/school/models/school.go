package models

import (
	commonModel "InitiaRe-website/internal/models"
)

type Response struct {
	Id         int    `json:"id"`
	SchoolName string `json:"school_name"`
	Note       string `json:"note"`
	StatusId   int    `json:"status_id"`
}

type RequestList struct {
	commonModel.RequestPaging
	SchoolName string `json:"school_name"`
	Note       string `json:note`
	StatusId   int    `json:"status_id"`
}

type ListPaging struct {
	commonModel.ListPaging
	Records []*Response
}

func (r *RequestList) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"status_id": r.StatusId,
		"page":      r.Page,
		"size":      r.Size,
		"sort_by":   r.SortBy,
		"order_by":  r.OrderBy,
	}
}
