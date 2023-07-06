package models

type UserInfo struct {
	Id                 int       `gorm:"primarykey;column:id" json:"id" redis:"id"`
	UserId 			   int 		 `gorm:"foreignkey;column:user_id" json:"user_id" redis:"user_id"`
	NumberUploaded     int 		 `gorm:"column:number_uploaded" json:"number_uploaded" redis:"number_uploaded" validate:"omitempty"`
	NumberPeerReviewed int 		 `gorm:"column:number_peer_reviewed" json:"number_peer_reviewed" redis:"number_peer_reviewed" validate:"omitempty"`
	NumberSpecReviewed int 		 `gorm:"column:number_spec_reviewed" json:"number_spec_reviewed" redis:"number_spec_reviewed" validate:"omitempty"`
}
