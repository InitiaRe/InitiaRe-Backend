package usecase

import (
	"context"

	"InitiaRe-website/internal/storage/models"
)

type IUseCase interface {
	UploadMedia(ctx context.Context, userId int, params *models.UploadRequest) (*models.Response, error)
}
