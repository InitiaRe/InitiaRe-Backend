package usecase

import (
	"context"

	"InitiaRe-website/constant"
	"InitiaRe-website/internal/article/entity"
	"InitiaRe-website/internal/article/models"
	articleCategoryModel "InitiaRe-website/internal/article_category/models"
	"InitiaRe-website/pkg/utils"

	"github.com/rs/zerolog/log"
)

func (u *usecase) Create(ctx context.Context, userId int, params *models.SaveRequest) (*models.Response, error) {
	article := &entity.Article{}
	article.ParseForCreate(params, userId)
	res, err := u.repo.Create(ctx, article)
	if err != nil {
		log.Error().Err(err).Str("service", "usecase.repo.Create").Send()
		return nil, utils.NewError(constant.STATUS_CODE_INTERNAL_SERVER, "Error when create article")
	}

	// Add primary category
	if params.CategoryId != 0 {
		if _, err := u.articleCategoryUc.Create(ctx, userId, &articleCategoryModel.SaveRequest{
			CategoryId: params.CategoryId,
			ArticleId:  res.Id,
			Type:       constant.ARTICLE_CATEGORY_TYPE_PRIMARY,
		}); err != nil {
			log.Error().Err(err).Str("service", "usecase.articleCategoryUc.Create").Send()
			return nil, utils.NewError(constant.STATUS_CODE_INTERNAL_SERVER, "Error when create article category")
		}
	}

	// Add secondary categories
	if len(params.SubCategoryIds) > 0 {
		for _, v := range params.SubCategoryIds {
			if _, err := u.articleCategoryUc.Create(ctx, userId, &articleCategoryModel.SaveRequest{
				CategoryId: v,
				ArticleId:  res.Id,
				Type:       constant.ARTICLE_CATEGORY_TYPE_SECONDARY,
			}); err != nil {
				log.Error().Err(err).Str("service", "usecase.articleCategoryUc.Create").Send()
				return nil, utils.NewError(constant.STATUS_CODE_INTERNAL_SERVER, "Error when create article category")
			}
		}
	}

	return res.Export(), nil
}
