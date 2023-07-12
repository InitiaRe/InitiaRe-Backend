package models

import (
	"time"	
)

type Articles struct {
	Id                 int       `gorm:"primarykey;column:id" json:"id" redis:"id"`
	UserId 			   int 		 `gorm:"column:user_id" json:"user_id" redis:"user_id"`
	CategoryId 		   int 		 `gorm:"column:category_id" json:"category_id" redis:"category_id"`
	StatusId 		   int 		 `gorm:"column:status_id" json:"status_id" redis:"status_id"`
	Content     	   string    `gorm:"column:content" json:"content" redis:"content"`			   
	PublishDate 	   time.Time `gorm:"column:publish_date" json:"publish_date,omitempty" redis:"publish_date"`
}
