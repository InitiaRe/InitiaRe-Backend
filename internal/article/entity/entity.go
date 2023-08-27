package entity

import (
	"time"

	"github.com/Ho-Minh/InitiaRe-website/constant"
	"github.com/Ho-Minh/InitiaRe-website/internal/article/models"
	"github.com/jinzhu/copier"
)

type Article struct {
	Id                int       `gorm:"primarykey;column:id" json:"id"`
	CategoryId        int       `gorm:"column:category_id" json:"category_id"`
	StatusId          int       `gorm:"column:status_id" json:"status_id"`
	Title             string    `gorm:"column:title;default:(-)" json:"title"`
	ShortBrief        string    `gorm:"column:short_brief;default:(-)" json:"short_brief,omitempty"`
	Content           string    `gorm:"column:content;default:(-)" json:"content"`
	Thumbnail         string    `gorm:"column:thumbnail;default:(-)" json:"thumbnail,omitempty"`
	PrePublishContent string    `gorm:"column:pre_publish_content;default:(-)" json:"pre_publish_content,omitempty"`
	PublishDate       time.Time `gorm:"column:publish_date;default:(-)" json:"publish_date,omitempty"`
	CreatedBy         int       `gorm:"column:created_by" json:"created_by"`
	CreatedAt         time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedBy         int       `gorm:"column:update_by;default:(-)" json:"update_by,omitempty"`
	UpdatedAt         time.Time `gorm:"autoUpdateTime;default:(-)" json:"updated_at,omitempty"`
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
