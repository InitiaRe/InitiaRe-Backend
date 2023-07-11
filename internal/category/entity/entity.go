package entity

import (
	"time"

	"github.com/Ho-Minh/InitiaRe-website/internal/category/models"
	"github.com/jinzhu/copier"
)

type Category struct {
	Id        int       `gorm:"primarykey;column:id" json:"id" redis:"id"`
	Category  string    `gorm:"column:category" json:"category" redis:"category"`
	CreatedBy int       `gorm:"column:created_by" json:"created_by,omitempty" redis:"created_by"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at,omitempty" redis:"created_at"`
	UpdatedBy int       `gorm:"column:update_by;default:(-)" json:"update_by,omitempty" redis:"update_by"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;default:(-)" json:"updated_at,omitempty" redis:"updated_at"`
}

func (c *Category) TableName() string {
	return "articles"
}

func (c *Category) Export() *models.Response {
	obj := &models.Response{}
	copier.Copy(obj, c) //nolint
	if !c.CreatedAt.IsZero() {
		obj.CreatedAt = c.CreatedAt.Format(time.RFC3339)
	}
	if !c.UpdatedAt.IsZero() {
		obj.UpdatedAt = c.UpdatedAt.Format(time.RFC3339)
	}
	return obj
}

func (c *Category) ExportList(in []*Category) []*models.Response {
	objs := make([]*models.Response, 0)
	for _, v := range in {
		objs = append(objs, v.Export())
	}
	return objs
}

func (c *Category) ParseFromSaveRequest(req *models.SaveRequest) {
	copier.Copy(c, req) //nolint
}

func (c *Category) ParseForCreate(req *models.SaveRequest, userId int) {
	c.ParseFromSaveRequest(req)
	c.CreatedBy = userId
}

func (c *Category) ParseForCreateMany(reqs []*models.SaveRequest, userId int) []*Category {
	objs := make([]*Category, 0)
	for _, v := range reqs {
		obj := &Category{}
		obj.ParseForCreate(v, userId)
		objs = append(objs, obj)
	}
	return objs
}

func (c *Category) ParseForUpdate(req *models.SaveRequest, userId int) {
	c.ParseFromSaveRequest(req)
	c.UpdatedBy = userId
}

func (c *Category) ParseForUpdateMany(reqs []*models.SaveRequest, userId int) []*Category {
	objs := make([]*Category, 0)
	for _, v := range reqs {
		obj := &Category{}
		obj.ParseForUpdate(v, userId)
		objs = append(objs, obj)
	}
	return objs
}
