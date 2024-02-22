package entity

import (
	"InitiaRe-website/constant"
	"InitiaRe-website/internal/user_info/models"

	"github.com/jinzhu/copier"
)

type UserInfo struct {
	Id                 int `gorm:"primarykey;column:id"`
	UserId             int `gorm:"column:user_id"`
	NumberUploaded     int `gorm:"column:number_uploaded"`
	NumberPeerReviewed int `gorm:"column:number_peer_reviewed"`
	NumberSpecReviewed int `gorm:"column:number_spec_reviewed"`
	Status             int `gorm:"column:status;default:(-)"`
	Role               int `gorm:"column:role;default:(-)"`
}

func (u *UserInfo) TableName() string {
	return "initiaRe_user_info"
}

func (u *UserInfo) Export() *models.Response {
	obj := &models.Response{}
	copier.Copy(obj, u) //nolint
	return obj
}

func (u *UserInfo) ExportList(in []*UserInfo) []*models.Response {
	objs := make([]*models.Response, 0)
	for _, v := range in {
		objs = append(objs, v.Export())
	}
	return objs
}

func (u *UserInfo) ParseFromSaveRequest(req *models.SaveRequest) {
	copier.Copy(u, req) //nolint
}

func (u *UserInfo) ParseForCreate(req *models.SaveRequest, userId int) {
	u.ParseFromSaveRequest(req)
	if u.Status == 0 {
		u.Status = constant.USER_STATUS_ACTIVE
	}
	if u.Role == 0 {
		u.Role = constant.USER_ROLE_NORMAL
	}
}

func (u *UserInfo) ParseForCreateMany(reqs []*models.SaveRequest, userId int) []*UserInfo {
	objs := make([]*UserInfo, 0)
	for _, v := range reqs {
		obj := &UserInfo{}
		obj.ParseForCreate(v, userId)
		objs = append(objs, obj)
	}
	return objs
}

func (u *UserInfo) ParseForUpdate(req *models.SaveRequest) {
	u.ParseFromSaveRequest(req)
}

func (u *UserInfo) ParseForUpdateMany(reqs []*models.SaveRequest) []*UserInfo {
	objs := make([]*UserInfo, 0)
	for _, v := range reqs {
		obj := &UserInfo{}
		obj.ParseForUpdate(v)
		objs = append(objs, obj)
	}
	return objs
}
