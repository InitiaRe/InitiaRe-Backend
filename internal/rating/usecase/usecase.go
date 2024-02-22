package usecase

import (
	"context"

	"InitiaRe-website/config"
	"InitiaRe-website/constant"
	"InitiaRe-website/internal/rating/entity"
	"InitiaRe-website/internal/rating/models"
	"InitiaRe-website/internal/rating/repository"
	"InitiaRe-website/pkg/utils"

	"github.com/rs/zerolog/log"
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
	return u.repo.GetArticleRating(ctx, articleId)
}

func (u *usecase) upsert(ctx context.Context, articleId, userId int) (int, error) {
	res, err := u.repo.GetOne(ctx, map[string]interface{}{
		"article_id": articleId,
		"user_id":    userId,
	})
	if err != nil {
		log.Error().Err(err).Str("prefix", "Rating").Str("service", "usecase.repo.GetOne").Send()
		return 0, utils.NewError(constant.STATUS_CODE_INTERNAL_SERVER, "Error when get rating")
	}

	// create new record
	if res.Id == 0 {
		if _, err = u.repo.Create(ctx, &entity.Rating{
			ArticleId: articleId,
			UserId:    userId,
			Rating:    1,
		}); err != nil {
			log.Error().Err(err).Str("prefix", "Rating").Str("service", "usecase.repo.Create").Send()
			return 0, utils.NewError(constant.STATUS_CODE_INTERNAL_SERVER, "Error when create rating")
		}
		return 1, nil
	}

	// update record
	var status int
	if res.Status == constant.RATING_STATUS_ACTIVE {
		status = constant.RATING_STATUS_INACTIVE
	} else {
		status = constant.RATING_STATUS_ACTIVE
	}
	if _, err := u.repo.Update(ctx, &entity.Rating{
		Id:     res.Id,
		Status: status,
	}); err != nil {
		log.Error().Err(err).Str("prefix", "Rating").Str("service", "usecase.repo.Update").Send()
		return 0, utils.NewError(constant.STATUS_CODE_INTERNAL_SERVER, "Error when update rating")
	}
	return 1, nil
}
