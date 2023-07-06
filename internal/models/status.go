package models

type Status struct {
	Id                 int       `gorm:"primarykey;column:id" json:"id" redis:"id"`
	Category     	   string    `gorm:"column:category" json:"category" redis:"category"`	
	StatusName 		   string    `gorm:"column:status_name" json:"status_name" redis:"status_name"`
}
