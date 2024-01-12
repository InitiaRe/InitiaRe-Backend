package usecase

import (
	"context"

	"github.com/Ho-Minh/InitiaRe-website/internal/auth/models"
)

type IUseCase interface {
	Register(ctx context.Context, params *models.SaveRequest) (*models.Response, error)
	Login(ctx context.Context, params *models.LoginRequest) (*models.UserWithToken, error)
	GetOne(ctx context.Context, params *models.RequestList) (*models.Response, error)
	ResetPassword(ctx context.Context, params *models.ResetPasswordRequest) (*models.Response, error)
}
