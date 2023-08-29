package repository

import (
	"context"

	"github.com/Ho-Minh/InitiaRe-website/internal/user_info/entity"
)

type IRepository interface {
	Create(ctx context.Context, obj *entity.UserInfo) (*entity.UserInfo, error)
	CreateMany(ctx context.Context, objs []*entity.UserInfo) (int, error)
	Update(ctx context.Context, obj *entity.UserInfo) (*entity.UserInfo, error)
	UpdateMany(ctx context.Context, objs []*entity.UserInfo) (int, error)
	Delete(ctx context.Context, id int) (int, error)
	DeleteMany(ctx context.Context, ids []int) (int, error)
	GetById(ctx context.Context, id int) (*entity.UserInfo, error)
	GetOne(ctx context.Context, queries map[string]interface{}) (*entity.UserInfo, error)
	GetList(ctx context.Context, queries map[string]interface{}) ([]*entity.UserInfo, error)
	GetListPaging(ctx context.Context, queries map[string]interface{}) ([]*entity.UserInfo, error)
	Count(ctx context.Context, queries map[string]interface{}) (int, error)
}
