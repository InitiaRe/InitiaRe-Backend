package usecase

import (
	"context"

	"github.com/Ho-Minh/InitiaRe-website/constant"
	"github.com/Ho-Minh/InitiaRe-website/internal/article/entity"
	"github.com/Ho-Minh/InitiaRe-website/internal/article/models"
	"github.com/Ho-Minh/InitiaRe-website/internal/article/repository"

	articleCategoryUc "github.com/Ho-Minh/InitiaRe-website/internal/article_category/usecase"
	categoryModel "github.com/Ho-Minh/InitiaRe-website/internal/category/models"
	categoryUc "github.com/Ho-Minh/InitiaRe-website/internal/category/usecase"

	commonModel "github.com/Ho-Minh/InitiaRe-website/internal/models"
	"github.com/Ho-Minh/InitiaRe-website/pkg/utils"
	"github.com/rs/zerolog/log"
)

type usecase struct {
	repo              repository.IRepository
	categoryUc        categoryUc.IUseCase
	articleCategoryUc articleCategoryUc.IUseCase
}

func InitUsecase(
	repo repository.IRepository,
	categoryUc categoryUc.IUseCase,
	articleCategoryUc articleCategoryUc.IUseCase,
) IUseCase {
	return &usecase{
		repo:              repo,
		categoryUc:        categoryUc,
		articleCategoryUc: articleCategoryUc,
	}
}

func (u *usecase) GetById(ctx context.Context, id int) (*models.Response, error) {
	record, err := u.repo.GetById(ctx, id)
	if err != nil {
		log.Error().Err(err).Str("prefix", "Article").Str("service", "usecase.repo.GetById").Send()
		return nil, utils.NewError(constant.STATUS_CODE_INTERNAL_SERVER, "Error when get article")
	}
	if record.Id == 0 {
		return nil, utils.NewError(constant.STATUS_CODE_NOT_FOUND, "Article not found")
	}

	res := record.Export()
	res.SubCategories, err = u.categoryUc.GetList(ctx, &categoryModel.RequestList{
		ArticleId: id,
	})
	if err != nil {
		log.Error().Err(err).Str("prefix", "Article").Str("service", "usecase.categoryUc.GetList").Send()
		return nil, utils.NewError(constant.STATUS_CODE_INTERNAL_SERVER, "Error when get article")
	}
	return res, nil
}

func (u *usecase) GetApprovedArticle(ctx context.Context) (*models.ApprovedList, error) {
	queries := map[string]interface{}{
		"status_id": constant.ARTICLE_STATUS_APPROVED,
	}

	record, err := u.repo.GetList(ctx, queries)
	if err != nil {
		log.Error().Err(err).Str("service", "usecase.repo.GetList").Send()
		return nil, utils.NewError(constant.STATUS_CODE_INTERNAL_SERVER, "Error when get approved article list")
	}
	count, err := u.repo.Count(ctx, queries)
	if err != nil {
		log.Error().Err(err).Str("service", "usecase.repo.Count").Send()
		return nil, utils.NewError(constant.STATUS_CODE_INTERNAL_SERVER, "Error when get approved article list")
	}

	return &models.ApprovedList{
		Total:   count,
		Records: (&entity.Article{}).ExportList(record),
	}, nil
}

func (u *usecase) GetList(ctx context.Context, params *models.RequestList) ([]*models.Response, error) {
	return nil, nil
}

func (u *usecase) GetListPaging(ctx context.Context, params *models.RequestList) (*models.ListPaging, error) {
	queries := params.ToMap()
	records, err := u.repo.GetListPaging(ctx, queries)
	if err != nil {
		log.Error().Err(err).Str("service", "usecase.repo.GetListPaging").Send()
		return nil, utils.NewError(constant.STATUS_CODE_INTERNAL_SERVER, "Error when get list article")
	}
	count, err := u.repo.Count(ctx, queries)
	if err != nil {
		log.Error().Err(err).Str("service", "usecase.repo.Count").Send()
		return nil, utils.NewError(constant.STATUS_CODE_INTERNAL_SERVER, "Error when get list article")
	}

	return &models.ListPaging{
		ListPaging: commonModel.ListPaging{
			Page:  params.Page,
			Size:  params.Size,
			Total: count,
		},
		Records: (&entity.Article{}).ExportList(records),
	}, nil
}

func (u *usecase) GetOne(ctx context.Context, params *models.RequestList) (*models.Response, error) {
	return nil, nil
}

func (u *usecase) CreateMany(ctx context.Context, userId int, params []*models.SaveRequest) (int, error) {
	return 0, nil
}

func (u *usecase) Update(ctx context.Context, userId int, params *models.SaveRequest) (*models.Response, error) {
	article := entity.Article{}
	article.ParseForUpdate(params, userId)

	if err := u.validateBeforeUpdate(ctx, article.Id, userId); err != nil {
		return nil, err
	}

	res, err := u.repo.Update(ctx, &article)
	if err != nil {
		log.Error().Err(err).Str("prefix", "Article").Str("service", "usecase.repo.Update").Send()
		return nil, utils.NewError(constant.STATUS_CODE_INTERNAL_SERVER, "Error when update article")
	}

	return res.Export(), err
}

func (u *usecase) UpdateMany(ctx context.Context, userId int, params []*models.SaveRequest) (int, error) {
	return 0, nil
}

func (u *usecase) Delete(ctx context.Context, id int) (int, error) {
	return 0, nil
}

func (u *usecase) DeleteMany(ctx context.Context, ids []int) (int, error) {
	return 0, nil
}

func (u *usecase) Approve(ctx context.Context, id int) error {
	article, err := u.repo.GetById(ctx, id)
	if err != nil {
		log.Error().Err(err).Str("prefix", "Article").Str("service", "usecase.repo.GetOne").Send()
		return utils.NewError(constant.STATUS_CODE_INTERNAL_SERVER, "Error when get article")
	}
	if article == nil {
		log.Error().Str("prefix", "Article").Msgf("Article not found with Id: %v", id)
		return utils.NewError(constant.STATUS_CODE_NOT_FOUND, "Article not found")
	}
	log.Info().Str("prefix", "Article").Msgf("Article status: %v && Article status pending: %v", article.StatusId, constant.ARTICLE_STATUS_PENDING)
	if article.StatusId != constant.ARTICLE_STATUS_PENDING {
		log.Error().Str("prefix", "Article").Msgf("Article is not pending with Id: %v", id)
		return utils.NewError(constant.STATUS_CODE_BAD_REQUEST, "Article is not pending")
	}
	article.StatusId = constant.ARTICLE_STATUS_APPROVED

	if _, err := u.repo.Update(ctx, article); err != nil {
		log.Error().Err(err).Str("prefix", "Article").Str("service", "usecase.repo.Update").Send()
		return utils.NewError(constant.STATUS_CODE_INTERNAL_SERVER, "Error when update article")
	}
	return nil
}

func (u *usecase) Disable(ctx context.Context, id int) error {
	article, err := u.repo.GetById(ctx, id)
	if err != nil {
		log.Error().Err(err).Str("prefix", "Article").Str("service", "usecase.repo.GetOne").Send()
		return utils.NewError(constant.STATUS_CODE_INTERNAL_SERVER, "Error when get article")
	}
	if article == nil {
		log.Error().Str("prefix", "Article").Msgf("Article not found with Id: %v", id)
		return utils.NewError(constant.STATUS_CODE_NOT_FOUND, "Article not found")
	}
	if article.StatusId == constant.ARTICLE_STATUS_HIDDEN {
		log.Error().Str("prefix", "Article").Msgf("Article is already disabled with Id: %v", id)
		return utils.NewError(constant.STATUS_CODE_BAD_REQUEST, "Article is already disabled ")
	}
	article.StatusId = constant.ARTICLE_STATUS_HIDDEN

	if _, err := u.repo.Update(ctx, article); err != nil {
		log.Error().Err(err).Str("prefix", "Article").Str("service", "usecase.repo.Update").Send()
		return utils.NewError(constant.STATUS_CODE_INTERNAL_SERVER, "Error when update article")
	}
	return nil
}

func (u *usecase) validateBeforeUpdate(ctx context.Context, id int, userId int) error {
	record, err := u.GetById(ctx, id)
	if err != nil {
		return err
	}

	if record.CreatedBy != userId {
		return utils.NewError(constant.STATUS_CODE_UNAUTHORIZED, "You do not have the permission to edit this document")
	}
	return nil
}
