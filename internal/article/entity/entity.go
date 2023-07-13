package entity

import (
	"time"

	"github.com/Ho-Minh/InitiaRe-website/internal/article/models"
	"github.com/jinzhu/copier"
)

type Article struct {
	Id          int       `gorm:"primarykey;column:id" json:"id" redis:"id"`
	CategoryId  int       `gorm:"column:category_id" json:"category_id" redis:"category_id"`
	StatusId    int       `gorm:"column:status_id" json:"status_id" redis:"status_id"`
	Content     string    `gorm:"column:content" json:"content" redis:"content"`
	PublishDate time.Time `gorm:"column:publish_date" json:"publish_date,omitempty" redis:"publish_date"`
	CreatedBy   int       `gorm:"column:created_by" json:"created_by,omitempty" redis:"created_by"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at,omitempty" redis:"created_at"`
	UpdatedBy   int       `gorm:"column:update_by;default:(-)" json:"update_by,omitempty" redis:"update_by"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime;default:(-)" json:"updated_at,omitempty" redis:"updated_at"`
}

func (a *Article) TableName() string {
	return "articles"
}

func (a *Article) Export() *models.Response {
	obj := &models.Response{}
	copier.Copy(obj, a) //nolint
	if !a.CreatedAt.IsZero() {
		obj.CreatedAt = a.CreatedAt.Format(time.RFC3339)
	}
	if !a.UpdatedAt.IsZero() {
		obj.UpdatedAt = a.UpdatedAt.Format(time.RFC3339)
	}
	return obj
}

func (a *Article) ExportList(in []*Article) []*models.Response {
	objs := make([]*models.Response, 0)
	for _, v := range in {
		objs = append(objs, v.Export())
	}
	return objs
}

func (a *Article) ParseFromSaveRequest(req *models.SaveRequest) {
	copier.Copy(a, req) //nolint
}

func (a *Article) ParseForCreate(req *models.SaveRequest, userId int) {
	a.ParseFromSaveRequest(req)
	a.CreatedBy = userId
}

func (a *Article) ParseForCreateMany(reqs []*models.SaveRequest, userId int) []*Article {
	objs := make([]*Article, 0)
	for _, v := range reqs {
		obj := &Article{}
		obj.ParseForCreate(v, userId)
		objs = append(objs, obj)
	}
	return objs
}

func (a *Article) ParseForUpdate(req *models.SaveRequest, userId int) {
	a.ParseFromSaveRequest(req)
	a.UpdatedBy = userId
}

func (a *Article) ParseForUpdateMany(reqs []*models.SaveRequest, userId int) []*Article {
	objs := make([]*Article, 0)
	for _, v := range reqs {
		obj := &Article{}
		obj.ParseForUpdate(v, userId)
		objs = append(objs, obj)
	}
	return objs
}
