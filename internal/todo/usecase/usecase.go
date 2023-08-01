package usecase

import (
	"context"

	"github.com/Ho-Minh/InitiaRe-website/constant"
	commonModel "github.com/Ho-Minh/InitiaRe-website/internal/models"
	"github.com/Ho-Minh/InitiaRe-website/internal/todo/entity"
	"github.com/Ho-Minh/InitiaRe-website/internal/todo/models"
	"github.com/Ho-Minh/InitiaRe-website/internal/todo/repository"
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
	record, err := u.repo.GetById(ctx, id)
	if err != nil {
		log.Error().Err(err).Str("prefix", "Todo").Str("service", "usecase.repo.GetById").Send()
		return nil, utils.NewError(constant.STATUS_CODE_INTERNAL_SERVER, "Error when get todo")
	}
	return record.Export(), nil
}

func (u *usecase) GetList(ctx context.Context, params *models.RequestList) ([]*models.Response, error) {
	queries := params.ToMap()
	records, err := u.repo.GetList(ctx, queries)
	if err != nil {
		log.Error().Err(err).Str("prefix", "Todo").Str("service", "usecase.repo.GetList").Send()
		return nil, utils.NewError(constant.STATUS_CODE_INTERNAL_SERVER, "Error when get list todo")
	}
	return (&entity.Todo{}).ExportList(records), nil
}

func (u *usecase) GetListPaging(ctx context.Context, params *models.RequestList) (*models.ListPaging, error) {
	queries := params.ToMap()
	records, err := u.repo.GetListPaging(ctx, queries)
	if err != nil {
		log.Error().Err(err).Str("prefix", "Todo").Str("service", "usecase.repo.GetListPaging").Send()
		return nil, utils.NewError(constant.STATUS_CODE_INTERNAL_SERVER, "Error when get list todo")
	}
	count, err := u.repo.Count(ctx, queries)
	if err != nil {
		log.Error().Err(err).Str("prefix", "Todo").Str("service", "usecase.repo.Count").Send()
		return nil, utils.NewError(constant.STATUS_CODE_INTERNAL_SERVER, "Error when get list todo")
	}

	return &models.ListPaging{
		ListPaging: commonModel.ListPaging{
			Page:  params.Page,
			Size:  params.Size,
			Total: count,
		},
		Records: (&entity.Todo{}).ExportList(records),
	}, nil
}

func (u *usecase) GetOne(ctx context.Context, params *models.RequestList) (*models.Response, error) {
	queries := params.ToMap()
	record, err := u.repo.GetOne(ctx, queries)
	if err != nil {
		log.Error().Err(err).Str("prefix", "Todo").Str("service", "usecase.repo.GetOne").Send()
		return nil, utils.NewError(constant.STATUS_CODE_INTERNAL_SERVER, "Error when get todo")
	}
	return record.Export(), nil
}

func (u *usecase) Create(ctx context.Context, userId int, params *models.SaveRequest) (*models.Response, error) {
	log.Info().Str("prefix", "Todo").Msgf("Create by user [%v] with params: %+v", userId, params)
	obj := &entity.Todo{}
	obj.ParseForCreate(params, userId)
	res, err := u.repo.Create(ctx, obj)
	if err != nil {
		log.Error().Err(err).Str("prefix", "Todo").Str("service", "usecase.repo.Create").Send()
		return nil, utils.NewError(constant.STATUS_CODE_INTERNAL_SERVER, "Error when create todo")
	}
	return res.Export(), nil
}

func (u *usecase) CreateMany(ctx context.Context, userId int, params []*models.SaveRequest) (int, error) {
	log.Info().Str("prefix", "Todo").Msgf("Create many by user [%v] with params: %+v", userId, params)
	objs := (&entity.Todo{}).ParseForCreateMany(params, userId)
	res, err := u.repo.CreateMany(ctx, objs)
	if err != nil {
		log.Error().Err(err).Str("prefix", "Todo").Str("service", "usecase.repo.CreateMany").Send()
		return 0, utils.NewError(constant.STATUS_CODE_INTERNAL_SERVER, "Error when create todo")
	}
	return res, nil
}

func (u *usecase) Update(ctx context.Context, userId int, params *models.SaveRequest) (*models.Response, error) {
	log.Info().Str("prefix", "Todo").Msgf("Update by user [%v] with params: %+v", userId, params)
	if err := u.validateBeforeUpdate(ctx, params.Id); err != nil {
		return nil, err
	}
	obj := &entity.Todo{}
	obj.ParseForUpdate(params, userId)
	res, err := u.repo.Update(ctx, obj)
	if err != nil {
		log.Error().Err(err).Str("prefix", "Todo").Str("service", "usecase.repo.Update").Send()
		return nil, utils.NewError(constant.STATUS_CODE_INTERNAL_SERVER, "Error when update todo")
	}
	return res.Export(), nil
}

func (u *usecase) UpdateMany(ctx context.Context, userId int, params []*models.SaveRequest) (int, error) {
	log.Info().Str("prefix", "Todo").Msgf("Update many by user [%v] with params: %+v", userId, params)
	for _, p := range params {
		if err := u.validateBeforeUpdate(ctx, p.Id); err != nil {
			return 0, err
		}
	}
	objs := (&entity.Todo{}).ParseForUpdateMany(params, userId)
	res, err := u.repo.UpdateMany(ctx, objs)
	if err != nil {
		log.Error().Err(err).Str("prefix", "Todo").Str("service", "usecase.repo.UpdateMany").Send()
		return 0, utils.NewError(constant.STATUS_CODE_INTERNAL_SERVER, "Error when update todo")
	}
	return res, nil
}

func (u *usecase) Delete(ctx context.Context, userId, id int) (int, error) {
	log.Info().Str("prefix", "Todo").Msgf("Delete by user [%v] with id: %v", userId, id)
	if err := u.validateBeforeUpdate(ctx, id); err != nil {
		return 0, err
	}
	res, err := u.repo.Delete(ctx, id)
	if err != nil {
		log.Error().Err(err).Str("prefix", "Todo").Str("service", "usecase.repo.Delete").Send()
		return 0, utils.NewError(constant.STATUS_CODE_INTERNAL_SERVER, "Error when delete todo")
	}
	return res, nil
}

func (u *usecase) DeleteMany(ctx context.Context, userId int, ids []int) (int, error) {
	log.Info().Str("prefix", "Todo").Msgf("Delete many by user [%v] with ids: %v", userId, ids)
	for _, id := range ids {
		if err := u.validateBeforeDelete(ctx, id); err != nil {
			return 0, err
		}
	}
	res, err := u.repo.DeleteMany(ctx, ids)
	if err != nil {
		log.Error().Err(err).Str("prefix", "Todo").Str("service", "usecase.repo.DeleteMany").Send()
		return 0, utils.NewError(constant.STATUS_CODE_INTERNAL_SERVER, "Error when delete todo")
	}
	return res, nil
}

func (u *usecase) validateBeforeUpdate(ctx context.Context, id int) error {
	record, err := u.repo.GetById(ctx, id)
	if err != nil {
		log.Error().Err(err).Str("prefix", "Todo").Str("service", "usecase.repo.GetById").Send()
		return utils.NewError(constant.STATUS_CODE_INTERNAL_SERVER, "Error when get todo")
	}
	if record.Id == 0 {
		return utils.NewError(constant.STATUS_CODE_NOT_FOUND, "Todo not found")
	}
	return nil
}

func (u *usecase) validateBeforeDelete(ctx context.Context, id int) error {
	record, err := u.repo.GetById(ctx, id)
	if err != nil {
		log.Error().Err(err).Str("prefix", "Todo").Str("service", "usecase.repo.GetById")
		return utils.NewError(constant.STATUS_CODE_INTERNAL_SERVER, "Error when get todo")
	}
	if record.Id == 0 {
		return utils.NewError(constant.STATUS_CODE_NOT_FOUND, "Todo not found")
	}
	return nil
}
