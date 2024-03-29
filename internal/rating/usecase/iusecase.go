package usecase

import (
	"context"

	"InitiaRe-website/internal/rating/models"
)

type IUseCase interface {
	Vote(ctx context.Context, params *models.SaveRequest) error
	GetRating(ctx context.Context, articleId int) (int, error)
}
