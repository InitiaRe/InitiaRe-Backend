package usecase

import (
	"context"

	"github.com/Ho-Minh/InitiaRe-website/internal/rating/models"
)

type IUseCase interface {
	Vote(ctx context.Context, params *models.SaveRequest) (*models.Response, error)
	GetRating(ctx context.Context, articleId int) (int, error)
}
