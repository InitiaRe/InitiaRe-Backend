package repository

import (
	"context"

	"InitiaRe-website/internal/storage/entity"
	"InitiaRe-website/internal/storage/models"
)

type IRepository interface {
	Create(ctx context.Context, obj *entity.Storage) (*entity.Storage, error)
	CreateMany(ctx context.Context, objs []*entity.Storage) ([]*entity.Storage, error)
	Update(ctx context.Context, obj *entity.Storage) (*entity.Storage, error)
	UpdateMany(ctx context.Context, objs []*entity.Storage) (int, error)
	Delete(ctx context.Context, id int) (int, error)
	DeleteMany(ctx context.Context, ids []int) (int, error)
	GetById(ctx context.Context, id int) (*entity.Storage, error)
	GetOne(ctx context.Context, queries map[string]interface{}) (*entity.Storage, error)
	GetList(ctx context.Context, queries map[string]interface{}) ([]*entity.Storage, error)
	GetListPaging(ctx context.Context, queries map[string]interface{}) ([]*entity.Storage, error)
	Count(ctx context.Context, queries map[string]interface{}) (int, error)
}

type IContainerRepository interface {
	Upload(ctx context.Context, req *models.UploadRequest) (string, error)
}
