package repository

import (
	"context"

	"github.com/Ho-Minh/InitiaRe-website/internal/rating/entity"
)

type IRepository interface {
	// CRUD
	Create(ctx context.Context, obj *entity.Rating) (*entity.Rating, error)
	CreateMany(ctx context.Context, objs []*entity.Rating) ([]*entity.Rating, error)
	Update(ctx context.Context, obj *entity.Rating) (*entity.Rating, error)
	UpdateMany(ctx context.Context, objs []*entity.Rating) (int, error)
	Delete(ctx context.Context, id int) (int, error)
	DeleteMany(ctx context.Context, ids []int) (int, error)
	GetById(ctx context.Context, id int) (*entity.Rating, error)
	GetOne(ctx context.Context, queries map[string]interface{}) (*entity.Rating, error)
	GetList(ctx context.Context, queries map[string]interface{}) ([]*entity.Rating, error)
	GetListPaging(ctx context.Context, queries map[string]interface{}) ([]*entity.Rating, error)
	Count(ctx context.Context, queries map[string]interface{}) (int, error)

	// Custom
	GetArticleRating(ctx context.Context, articleId int) (int, error)
}
