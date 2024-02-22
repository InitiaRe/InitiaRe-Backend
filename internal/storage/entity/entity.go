package entity

import (
	"time"

	"InitiaRe-website/internal/storage/models"

	"github.com/jinzhu/copier"
)

type Storage struct {
	Id          int       `gorm:"primarykey;column:id" json:"id" redis:"id"`
	DownloadUrl string    `gorm:"column:download_url" json:"download_url,omitempty" redis:"download_url"`
	Type        string    `gorm:"column:type" json:"type,omitempty" redis:"type"`
	Token       string    `gorm:"column:token;default:(-)" json:"token,omitempty" redis:"token"`
	LifeTime    int       `gorm:"column:life_time;default:(-)" json:"life_time,omitempty" redis:"life_time"`
	CreatedBy   int       `gorm:"column:created_by" json:"created_by,omitempty" redis:"created_by"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at,omitempty" redis:"created_at"`
}

func (s *Storage) TableName() string {
	return "initiaRe_storage"
}

func (s *Storage) Export() *models.Response {
	obj := &models.Response{}
	copier.Copy(obj, s) //nolint
	if !s.CreatedAt.IsZero() {
		obj.CreatedAt = s.CreatedAt.Format(time.RFC3339)
	}
	return obj
}

func (s *Storage) ExportList(in []*Storage) []*models.Response {
	objs := make([]*models.Response, 0)
	for _, v := range in {
		objs = append(objs, v.Export())
	}
	return objs
}

func (s *Storage) ParseFromSaveRequest(req *models.SaveRequest) {
	copier.Copy(s, req) //nolint
}

func (s *Storage) ParseForCreate(req *models.SaveRequest, userId int) {
	s.ParseFromSaveRequest(req)
	s.CreatedBy = userId
}

func (s *Storage) ParseForCreateMany(reqs []*models.SaveRequest, userId int) []*Storage {
	objs := make([]*Storage, 0)
	for _, v := range reqs {
		obj := &Storage{}
		obj.ParseForCreate(v, userId)
		objs = append(objs, obj)
	}
	return objs
}
