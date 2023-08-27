package models

type Status struct {
	StatusId   int    `gorm:"column:status_id" json:"id"`
	Category   string `gorm:"column:category" json:"category"`
	StatusName string `gorm:"column:status_name" json:"status_name"`
}

func (s *Status) TableName() string {
	return "initiaRe_status"
}
