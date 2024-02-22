package usecase

import (
	"context"

	"InitiaRe-website/constant"
	"InitiaRe-website/internal/category/entity"
	"InitiaRe-website/internal/category/models"
	"InitiaRe-website/internal/category/repository"
	commonModel "InitiaRe-website/internal/models"
	"InitiaRe-website/pkg/utils"

	"github.com/rs/zerolog/log"
)

type usecase struct {
	repo repository.IRepository
}

func InitUsecase(repo repository.IRepository) IUseCase {
	return &usecase{
		repo: repo,
	}
}

func (u *usecase) GetById(ctx context.Context, id int) (*models.Response, error) {
	record, err := u.repo.GetById(ctx, id)
	if err != nil {
		log.Error().Err(err).Str("prefix", "Category").Str("service", "usecase.repo.GetById").Send()
		return nil, utils.NewError(constant.STATUS_CODE_INTERNAL_SERVER, "Error when get category")
	}

	if record.Id == 0 {
		return nil, utils.NewError(constant.STATUS_CODE_NOT_FOUND, "Category not found")
	}

	return record.Export(), nil
}

func (u *usecase) GetList(ctx context.Context, params *models.RequestList) ([]*models.Response, error) {
	queries := params.ToMap()
	records, err := u.repo.GetList(ctx, queries)
	if err != nil {
		log.Error().Err(err).Str("prefix", "Category").Str("service", "usecase.repo.GetList").Send()
		return nil, utils.NewError(constant.STATUS_CODE_INTERNAL_SERVER, "Error when get list category")
	}
	return (&entity.Category{}).ExportList(records), nil
}

func (u *usecase) GetListPaging(ctx context.Context, params *models.RequestList) (*models.ListPaging, error) {
	queries := params.ToMap()
	records, err := u.repo.GetListPaging(ctx, queries)
	if err != nil {
		log.Error().Err(err).Str("prefix", "Category").Str("service", "usecase.repo.GetListPaging").Send()
		return nil, utils.NewError(constant.STATUS_CODE_INTERNAL_SERVER, "Error when get list category")
	}
	count, err := u.repo.Count(ctx, queries)
	if err != nil {
		log.Error().Err(err).Str("prefix", "Category").Str("service", "usecase.repo.Count").Send()
		return nil, utils.NewError(constant.STATUS_CODE_INTERNAL_SERVER, "Error when get list category")
	}

	return &models.ListPaging{
		ListPaging: commonModel.ListPaging{
			Page:  params.Page,
			Size:  params.Size,
			Total: count,
		},
		Records: (&entity.Category{}).ExportList(records),
	}, nil
}

func (u *usecase) GetOne(ctx context.Context, params *models.RequestList) (*models.Response, error) {
	return nil, nil
}

func (u *usecase) Create(ctx context.Context, userId int, params *models.SaveRequest) (*models.Response, error) {
	obj := &entity.Category{}
	obj.ParseForCreate(params, userId)
	res, err := u.repo.Create(ctx, obj)
	if err != nil {
		log.Error().Err(err).Str("prefix", "Category").Str("service", "usecase.repo.Create").Send()
		return nil, utils.NewError(constant.STATUS_CODE_INTERNAL_SERVER, "Error when create category")
	}
	return res.Export(), nil
}

func (u *usecase) CreateMany(ctx context.Context, userId int, params []*models.SaveRequest) (int, error) {
	return 0, nil
}

func (u *usecase) Update(ctx context.Context, userId int, params *models.SaveRequest) (*models.Response, error) {
	log.Info().Str("prefix", "Category").Msgf("Update by user [%v] with params: %+v", userId, params)
	if err := u.validateBeforeUpdate(ctx, params.Id); err != nil {
		return nil, err
	}
	obj := &entity.Category{}
	obj.ParseForUpdate(params, userId)
	res, err := u.repo.Update(ctx, obj)
	if err != nil {
		log.Error().Err(err).Str("prefix", "Category").Str("service", "usecase.repo.Update").Send()
		return nil, utils.NewError(constant.STATUS_CODE_INTERNAL_SERVER, "Error when update category")
	}
	return res.Export(), nil
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

func (u *usecase) validateBeforeUpdate(ctx context.Context, id int) error {
	if _, err := u.GetById(ctx, id); err != nil {
		return err
	}
	return nil
}
