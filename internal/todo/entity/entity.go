package entity

import (
	"time"

	"github.com/Ho-Minh/InitiaRe-website/internal/todo/models"

	"github.com/jinzhu/copier"
)

type Todo struct {
	Id        int       `gorm:"primarykey;column:id" json:"id" redis:"id"`
	Content   string    `gorm:"column:content" json:"content,omitempty" redis:"content"`
	CreatedBy int       `gorm:"column:created_by" json:"created_by,omitempty" redis:"created_by"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at,omitempty" redis:"created_at"`
	UpdatedBy int       `gorm:"column:update_by;default:(-)" json:"update_by,omitempty" redis:"update_by"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;default:(-)" json:"updated_at,omitempty" redis:"updated_at"`
}

func (t *Todo) TableName() string {
	return "initiaRe_todo"
}

func (a *Todo) Export() *models.Response {
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

func (a *Todo) ExportList(in []*Todo) []*models.Response {
	objs := make([]*models.Response, 0)
	for _, v := range in {
		objs = append(objs, v.Export())
	}
	return objs
}

func (a *Todo) ParseFromSaveRequest(req *models.SaveRequest) {
	copier.Copy(a, req) //nolint
}

func (a *Todo) ParseForCreate(req *models.SaveRequest, userId int) {
	a.ParseFromSaveRequest(req)
	a.CreatedBy = userId
}

func (a *Todo) ParseForCreateMany(reqs []*models.SaveRequest, userId int) []*Todo {
	objs := make([]*Todo, 0)
	for _, v := range reqs {
		obj := &Todo{}
		obj.ParseForCreate(v, userId)
		objs = append(objs, obj)
	}
	return objs
}

func (a *Todo) ParseForUpdate(req *models.SaveRequest, userId int) {
	a.ParseFromSaveRequest(req)
	a.UpdatedBy = userId
}

func (a *Todo) ParseForUpdateMany(reqs []*models.SaveRequest, userId int) []*Todo {
	objs := make([]*Todo, 0)
	for _, v := range reqs {
		obj := &Todo{}
		obj.ParseForUpdate(v, userId)
		objs = append(objs, obj)
	}
	return objs
}
