package usecase

import (
	"context"

	"github.com/Ho-Minh/InitiaRe-website/config"
	"github.com/Ho-Minh/InitiaRe-website/internal/rating/models"
	"github.com/Ho-Minh/InitiaRe-website/internal/rating/repository"
)

type usecase struct {
	cfg  *config.Config
	repo repository.IRepository
}

func InitUsecase(cfg *config.Config, repo repository.IRepository) IUseCase {
	return &usecase{
		cfg:  cfg,
		repo: repo,
	}
}

func (u *usecase) Vote(ctx context.Context, params *models.SaveRequest) (*models.Response, error) {
	return nil, nil
}
func (u *usecase) GetRating(ctx context.Context, articleId int) (int, error) {
	return 0, nil
}
