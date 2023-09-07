package entity

import (
	"github.com/Ho-Minh/InitiaRe-website/internal/article_category/models"

	"github.com/jinzhu/copier"
)

type ArticleCategory struct {
	Id         int `gorm:"primarykey;column:id"`
	CategoryId int `gorm:"column:category_id"`
	ArticleId  int `gorm:"column:article_id"`
	Type       int `gorm:"column:type"`
}

func (a *ArticleCategory) TableName() string {
	return "initiaRe_article_category"
}

func (a *ArticleCategory) Export() *models.Response {
	obj := &models.Response{}
	copier.Copy(obj, a) //nolint
	return obj
}

func (a *ArticleCategory) ExportList(in []*ArticleCategory) []*models.Response {
	objs := make([]*models.Response, 0)
	for _, v := range in {
		objs = append(objs, v.Export())
	}
	return objs
}

func (a *ArticleCategory) ParseFromSaveRequest(req *models.SaveRequest) {
	copier.Copy(a, req) //nolint
}

func (a *ArticleCategory) ParseForCreate(req *models.SaveRequest, userId int) {
	a.ParseFromSaveRequest(req)
}

func (a *ArticleCategory) ParseForCreateMany(reqs []*models.SaveRequest, userId int) []*ArticleCategory {
	objs := make([]*ArticleCategory, 0)
	for _, v := range reqs {
		obj := &ArticleCategory{}
		obj.ParseForCreate(v, userId)
		objs = append(objs, obj)
	}
	return objs
}

func (a *ArticleCategory) ParseForUpdate(req *models.SaveRequest, userId int) {
	a.ParseFromSaveRequest(req)
}

func (a *ArticleCategory) ParseForUpdateMany(reqs []*models.SaveRequest, userId int) []*ArticleCategory {
	objs := make([]*ArticleCategory, 0)
	for _, v := range reqs {
		obj := &ArticleCategory{}
		obj.ParseForUpdate(v, userId)
		objs = append(objs, obj)
	}
	return objs
}
