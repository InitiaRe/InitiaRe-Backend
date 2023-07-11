package repository

import (
	"context"

	"github.com/Ho-Minh/InitiaRe-website/internal/category/entity"
)

type IRepository interface {
	Create(ctx context.Context, obj *entity.Category) (*entity.Category, error)
	CreateMany(ctx context.Context, objs []*entity.Category) (int, error)
	Update(ctx context.Context, obj *entity.Category) (*entity.Category, error)
	UpdateMany(ctx context.Context, objs []*entity.Category) (int, error)
	GetById(ctx context.Context, id int) (*entity.Category, error)
	GetOne(ctx context.Context, queries map[string]interface{}) (*entity.Category, error)
	GetList(ctx context.Context, queries map[string]interface{}) ([]*entity.Category, error)
	GetListPaging(ctx context.Context, queries map[string]interface{}) ([]*entity.Category, error)
	Count(ctx context.Context, queries map[string]interface{}) (int, error)
}
