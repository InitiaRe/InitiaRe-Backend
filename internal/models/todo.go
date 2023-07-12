package models

import (
	"time"
)

type Todo struct {
	Id        int       `gorm:"primarykey;column:id" json:"id" redis:"id"`
	Content   string    `gorm:"column:content" json:"content,omitempty" redis:"content" validate:"required,lte=1024"`
	CreatedBy int       `gorm:"column:created_by" json:"created_by,omitempty" redis:"created_by" validate:"required"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at,omitempty" redis:"created_at" validate:"required"`
	UpdateBy  int       `gorm:"column:update_by;default:(-)" json:"update_by,omitempty" redis:"update_by"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;default:(-)" json:"updated_at,omitempty" redis:"updated_at"`
}
