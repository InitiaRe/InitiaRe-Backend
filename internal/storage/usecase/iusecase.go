package usecase

import (
	"context"

	"github.com/Ho-Minh/InitiaRe-website/internal/storage/models"
)

type IUseCase interface {
	UploadMedia(ctx context.Context, userId int, params *models.UploadRequest) (*models.Response, error)
}
