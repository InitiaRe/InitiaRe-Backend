package entity

import (
	"time"

	"InitiaRe-website/constant"
	"InitiaRe-website/internal/article/models"

	"github.com/jinzhu/copier"
)

type Article struct {
	Id                int       `gorm:"primarykey;column:id"`
	CategoryId        int       `gorm:"column:category_id"`
	StatusId          int       `gorm:"column:status_id"`
	Title             string    `gorm:"column:title;default:(-)"`
	ShortBrief        string    `gorm:"column:short_brief;default:(-)"`
	Content           string    `gorm:"column:content;default:(-)"`
	Thumbnail         string    `gorm:"column:thumbnail;default:(-)"`
	PrePublishContent string    `gorm:"column:pre_publish_content;default:(-)"`
	PublishDate       time.Time `gorm:"column:publish_date;default:(-)"`
	TypeId            int       `gorm:"column:type_id;default:(-)"` // 1: research, 2: review, 3: research proposal
	CreatedBy         int       `gorm:"column:created_by"`
	CreatedAt         time.Time `gorm:"autoCreateTime"`
	UpdatedBy         int       `gorm:"column:updated_by;default:(-)"`
	UpdatedAt         time.Time `gorm:"autoUpdateTime;default:(-)"`

	// Custom fields
	StatusName   string `gorm:"->;-:migration"`
	CategoryName string `gorm:"->;-:migration"`
	Email        string `gorm:"->;-:migration"`
}

func (a *Article) TableName() string {
	return "initiaRe_article"
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
	if a.TypeId != 0 {
		obj.TypeName = constant.ARTICLE_TYPE[a.TypeId]
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

func (a *Article) parseFromSaveRequest(req *models.SaveRequest) {
	copier.Copy(a, req) //nolint
}

func (a *Article) ParseForCreate(req *models.SaveRequest, userId int) {
	a.parseFromSaveRequest(req)
	a.CreatedBy = userId
	a.StatusId = constant.ARTICLE_STATUS_PENDING
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
	a.parseFromSaveRequest(req)
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
