package repository

import (
	"context"

	"github.com/Ho-Minh/InitiaRe-website/internal/auth/entity"
)

type IRepository interface {
	Create(ctx context.Context, todo *entity.User) (*entity.User, error)
	GetById(ctx context.Context, userId int) (*entity.User, error)
	GetOne(ctx context.Context, queries map[string]interface{}) (*entity.User, error)
}

type ICacheRepository interface {
	GetById(ctx context.Context, key string) (*entity.User, error)
	SetUser(ctx context.Context, key string, seconds int, user *entity.User) error
}
