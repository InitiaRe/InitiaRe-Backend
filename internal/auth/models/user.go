package models

import (
	"time"

	commonModel "github.com/Ho-Minh/InitiaRe-website/internal/models"
	"github.com/jinzhu/copier"

	"github.com/google/uuid"
)

type RequestList struct {
	commonModel.RequestPaging
	Email string
}

func (r *RequestList) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"email":    r.Email,
		"page":     r.Page,
		"size":     r.Size,
		"sort_by":  r.SortBy,
		"order_by": r.OrderBy,
	}
}

type Response struct {
	Id        int       `json:"id"`
	UserId    uuid.UUID `json:"user_id"`
	FirstName string    `json:"first_name,omitempty"`
	LastName  string    `json:"last_name,omitempty"`
	Email     string    `json:"email,omitempty"`
	School    string    `json:"school,omitempty"`
	Gender    string    `json:"gender,omitempty"`

	// Custom fields
	Status int `json:"status,omitempty"`
}

type SaveRequest struct {
	Id        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Gender    string    `json:"gender"`
	School    string    `json:"school"`
	Birthday  time.Time `json:"birthday"`
}

type LoginRequest struct {
	Email    string
	Password string
}

type RegisterRequest struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email,omitempty"`
	School    string `json:"school,omitempty"`
	Gender    string `json:"gender,omitempty"`
	Password  string `json:"password"`
}

func (r *RegisterRequest) ToSaveRequest() *SaveRequest {
	req := &SaveRequest{}
	copier.Copy(req, r)
	return req
}

// User sign in response
type UserWithToken struct {
	User  *Response `json:"user,omitempty"`
	Token string    `json:"token"`
}
