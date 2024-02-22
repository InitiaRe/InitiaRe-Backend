package repository

import (
	"context"

	"InitiaRe-website/internal/article_category/entity"
)

type IRepository interface {
	Create(ctx context.Context, obj *entity.ArticleCategory) (*entity.ArticleCategory, error)
	CreateMany(ctx context.Context, objs []*entity.ArticleCategory) ([]*entity.ArticleCategory, error)
	Update(ctx context.Context, obj *entity.ArticleCategory) (*entity.ArticleCategory, error)
	UpdateMany(ctx context.Context, objs []*entity.ArticleCategory) (int, error)
	Delete(ctx context.Context, id int) (int, error)
	DeleteMany(ctx context.Context, ids []int) (int, error)
	GetById(ctx context.Context, id int) (*entity.ArticleCategory, error)
	GetOne(ctx context.Context, queries map[string]interface{}) (*entity.ArticleCategory, error)
	GetList(ctx context.Context, queries map[string]interface{}) ([]*entity.ArticleCategory, error)
	GetListPaging(ctx context.Context, queries map[string]interface{}) ([]*entity.ArticleCategory, error)
	Count(ctx context.Context, queries map[string]interface{}) (int, error)
}
