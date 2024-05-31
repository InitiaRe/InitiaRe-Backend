package entity

import (
	"InitiaRe-website/internal/school/models"

	"github.com/jinzhu/copier"
)

type School struct {
	Id         int    `gorm:"primarykey;column:id"`
	SchoolName string `gorm:"column:school_name"`
	Note       string `gorm:"column:note;default:(-)"`
	StatusId   int    `gorm:"column:status_id"` // 1: Pending, 2: Rejected, 3: Verified

	// Custom fields
	StatusName string `gorm:"->;-:migration"`
}

func (a *School) TableName() string {
	return "initiaRe_school"
}

func (a *School) Export() *models.Response {
	obj := &models.Response{}
	copier.Copy(obj, a) //nolint
	return obj
}

func (a *School) ExportList(in []*School) []*models.Response {
	objs := make([]*models.Response, 0)
	for _, v := range in {
		objs = append(objs, v.Export())
	}
	return objs
}

// func (a *School) parseFromSaveRequest(req *models.SaveRequest) {
// 	copier.Copy(a, req) //nolint
// }
//
// func (a *School) ParseForCreate(req *models.SaveRequest, userId int) {
// 	a.parseFromSaveRequest(req)
// 	a.StatusId = constant.SCHOOL_STATUS_PENDING
// }
//
// func (a *School) ParseForCreateMany(reqs []*models.SaveRequest, userId int) []*School {
// 	objs := make([]*School, 0)
// 	for _, v := range reqs {
// 		obj := &School{}
// 		obj.ParseForCreate(v, userId)
// 		objs = append(objs, obj)
// 	}
// 	return objs
// }
//
// func (a *School) ParseForUpdate(req *models.SaveRequest, userId int) {
// 	a.parseFromSaveRequest(req)
// }
//
// func (a *School) ParseForUpdateMany(reqs []*models.SaveRequest, userId int) []*School {
// 	objs := make([]*School, 0)
// 	for _, v := range reqs {
// 		obj := &School{}
// 		obj.ParseForUpdate(v, userId)
// 		objs = append(objs, obj)
// 	}
// 	return objs
// }
