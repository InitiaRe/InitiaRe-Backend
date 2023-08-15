package models

type UserInfo struct {
	Id                 int       `gorm:"primarykey;column:id" json:"id" redis:"id"`
	UserId 			   int 		 `gorm:"column:user_id" json:"user_id" redis:"user_id"`
	NumberUploaded     int 		 `gorm:"column:number_uploaded" json:"number_uploaded" redis:"number_uploaded"`
	NumberPeerReviewed int 		 `gorm:"column:number_peer_reviewed" json:"number_peer_reviewed" redis:"number_peer_reviewed"`
	NumberSpecReviewed int 		 `gorm:"column:number_spec_reviewed" json:"number_spec_reviewed" redis:"number_spec_reviewed"`
}


func (u *UserInfo) TableName() string {
	return "initiaRe_user_info"
}