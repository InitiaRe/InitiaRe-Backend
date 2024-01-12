package entity

import (
	"github.com/Ho-Minh/InitiaRe-website/constant"
	"github.com/Ho-Minh/InitiaRe-website/internal/rating/models"
	"github.com/jinzhu/copier"
)

type Rating struct {
	Id        int `gorm:"primarykey;column:id"`
	UserId    int `gorm:"column:user_id"`
	ArticleId int `gorm:"column:article_id"`
	Rating    int `gorm:"column:rating"`
	Status    int `gorm:"column:status"`
}

func (r *Rating) TableName() string {
	return "initiaRe_rating"
}

func (a *Rating) Export() *models.Response {
	obj := &models.Response{}
	copier.Copy(obj, a) //nolint
	return obj
}

func (a *Rating) ExportList(in []*Rating) []*models.Response {
	objs := make([]*models.Response, 0)
	for _, v := range in {
		objs = append(objs, v.Export())
	}
	return objs
}

func (a *Rating) parseFromSaveRequest(req *models.SaveRequest) {
	copier.Copy(a, req) //nolint
}

func (a *Rating) ParseForCreate(req *models.SaveRequest, userId int) {
	a.parseFromSaveRequest(req)
	if req.Status == 0 {
		a.Status = constant.RATING_STATUS_ACTIVE
	}
}

func (a *Rating) ParseForCreateMany(reqs []*models.SaveRequest, userId int) []*Rating {
	objs := make([]*Rating, 0)
	for _, v := range reqs {
		obj := &Rating{}
		obj.ParseForCreate(v, userId)
		objs = append(objs, obj)
	}
	return objs
}

func (a *Rating) ParseForUpdate(req *models.SaveRequest, userId int) {
	a.parseFromSaveRequest(req)
}

func (a *Rating) ParseForUpdateMany(reqs []*models.SaveRequest, userId int) []*Rating {
	objs := make([]*Rating, 0)
	for _, v := range reqs {
		obj := &Rating{}
		obj.ParseForUpdate(v, userId)
		objs = append(objs, obj)
	}
	return objs
}
