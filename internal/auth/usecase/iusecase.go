package usecase

import (
	"context"

	"github.com/Ho-Minh/InitiaRe-website/internal/auth/models"
)

type IUseCase interface {
	Register(ctx context.Context, params *models.SaveRequest) (*models.Response, error)
	Login(ctx context.Context, params *models.LoginRequest) (*models.UserWithToken, error)
}
