package repository

import (
	"context"

	"github.com/Ho-Minh/InitiaRe-website/internal/article/entity"
)

type IRepository interface {
	Create(ctx context.Context, obj *entity.Article) (*entity.Article, error)
	CreateMany(ctx context.Context, objs []*entity.Article) (int, error)
	Update(ctx context.Context, obj *entity.Article) (*entity.Article, error)
	UpdateMany(ctx context.Context, objs []*entity.Article) (int, error)
	Delete(ctx context.Context, id int) (int, error)
	DeleteMany(ctx context.Context, ids []int) (int, error)
	GetById(ctx context.Context, id int) (*entity.Article, error)
	GetOne(ctx context.Context, queries map[string]interface{}) (*entity.Article, error)
	GetList(ctx context.Context, queries map[string]interface{}) ([]*entity.Article, error)
	GetListPaging(ctx context.Context, queries map[string]interface{}) ([]*entity.Article, error)
	Count(ctx context.Context, queries map[string]interface{}) (int, error)
}
