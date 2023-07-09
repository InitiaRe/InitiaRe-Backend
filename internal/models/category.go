package models

type Category struct {
	Id                 int       `gorm:"primarykey;column:id" json:"id" redis:"id"`
	Category     	   string    `gorm:"column:category" json:"category" redis:"category"`			   
}
