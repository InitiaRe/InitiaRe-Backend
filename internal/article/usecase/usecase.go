package usecase

import (
	"context"

	"github.com/Ho-Minh/InitiaRe-website/constant"
	"github.com/Ho-Minh/InitiaRe-website/internal/article/entity"
	"github.com/Ho-Minh/InitiaRe-website/internal/article/models"
	"github.com/Ho-Minh/InitiaRe-website/internal/article/repository"

	commonModel "github.com/Ho-Minh/InitiaRe-website/internal/models"
	"github.com/Ho-Minh/InitiaRe-website/pkg/utils"
	"github.com/rs/zerolog/log"
)

type usecase struct {
	repo repository.IRepository
}

func NewUseCase(repo repository.IRepository) IUseCase {
	return &usecase{
		repo: repo,
	}
}

func (u *usecase) GetById(ctx context.Context, id int) (*models.Response, error) {
	return nil, nil
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

func (u *usecase) Create(ctx context.Context, userId int, params *models.SaveRequest) (*models.Response, error) {
	article := &entity.Article{}
	article.ParseForCreate(params, userId)
	res, err := u.repo.Create(ctx, article)
	if err != nil {
		log.Error().Err(err).Str("service", "usecase.repo.Create").Send()
		return nil, utils.NewError(constant.STATUS_CODE_INTERNAL_SERVER, "Error when create article")
	}

	return res.Export(), nil
}

func (u *usecase) CreateMany(ctx context.Context, userId int, params []*models.SaveRequest) (int, error) {
	return 0, nil
}

func (u *usecase) Update(ctx context.Context, userId int, params *models.SaveRequest) (*models.Response, error) {
	article := entity.Article{}
	article.ParseForUpdate(params, userId)

	// validation
	record, err := u.repo.GetById(ctx, article.Id)
	if err != nil {
		log.Error().Err(err).Str("service", "usecase.repo.GetById").Send()
		return nil, utils.NewError(constant.STATUS_CODE_INTERNAL_SERVER, "Error when get article")
	}
	if record.CreatedBy != userId {
		return nil, utils.NewError(constant.STATUS_CODE_INTERNAL_SERVER, "Error when update article")
	}

	res, err := u.repo.Update(ctx, &article)
	if err != nil {
		log.Error().Err(err).Str("service", "usecase.repo.Update").Send()
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
