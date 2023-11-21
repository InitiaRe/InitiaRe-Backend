package usecase

import (
	"context"
)

type IUseCase interface {
	Enable(ctx context.Context, userId int) error
	Disable(ctx context.Context, userId int) error
	IsAdmin(ctx context.Context, userId int) (bool, error)
	PromoteAdmin(ctx context.Context, userId int, email string) error
}
