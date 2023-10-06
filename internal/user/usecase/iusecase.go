package usecase

import (
	"context"
)

type IUseCase interface {
	Enable(ctx context.Context, userId int) error
	Disable(ctx context.Context, userId int) error
}
