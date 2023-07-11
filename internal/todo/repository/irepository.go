package repository

import (
	"context"

	"github.com/Ho-Minh/InitiaRe-website/internal/todo/entity"
)

type IRepository interface {
	Create(ctx context.Context, todo *entity.Todo) (*entity.Todo, error)
	CreateMany(ctx context.Context, objs []*entity.Todo) (int, error)
	Update(ctx context.Context, obj *entity.Todo) (*entity.Todo, error)
	UpdateMany(ctx context.Context, objs []*entity.Todo) (int, error)
	GetById(ctx context.Context, id int) (*entity.Todo, error)
	GetOne(ctx context.Context, queries map[string]interface{}) (*entity.Todo, error)
	GetList(ctx context.Context, queries map[string]interface{}) ([]*entity.Todo, error)
	GetListPaging(ctx context.Context, queries map[string]interface{}) ([]*entity.Todo, error)
	Count(ctx context.Context, queries map[string]interface{}) (int, error)
}
