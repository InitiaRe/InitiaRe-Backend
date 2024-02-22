package repository

import (
	"context"

	"InitiaRe-website/internal/category/entity"
)

type IRepository interface {
	Create(ctx context.Context, obj *entity.Category) (*entity.Category, error)
	CreateMany(ctx context.Context, objs []*entity.Category) ([]*entity.Category, error)
	Update(ctx context.Context, obj *entity.Category) (*entity.Category, error)
	UpdateMany(ctx context.Context, objs []*entity.Category) (int, error)
	Delete(ctx context.Context, id int) (int, error)
	DeleteMany(ctx context.Context, ids []int) (int, error)
	GetById(ctx context.Context, id int) (*entity.Category, error)
	GetOne(ctx context.Context, queries map[string]interface{}) (*entity.Category, error)
	GetList(ctx context.Context, queries map[string]interface{}) ([]*entity.Category, error)
	GetListPaging(ctx context.Context, queries map[string]interface{}) ([]*entity.Category, error)
	Count(ctx context.Context, queries map[string]interface{}) (int, error)
}
