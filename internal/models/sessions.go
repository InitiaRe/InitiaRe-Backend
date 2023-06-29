package models

import "github.com/google/uuid"

// Session model
type Session struct {
	SessionId string    `json:"session_id" redis:"session_id"`
	UserId    uuid.UUID `json:"user_id" redis:"user_id"`
}
