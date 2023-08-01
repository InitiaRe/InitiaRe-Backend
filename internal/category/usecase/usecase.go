package usecase

import (
	"context"

	"github.com/Ho-Minh/InitiaRe-website/internal/category/entity"
	"github.com/Ho-Minh/InitiaRe-website/internal/category/models"
	"github.com/Ho-Minh/InitiaRe-website/internal/category/repository"
	"github.com/Ho-Minh/InitiaRe-website/internal/constants"
	commonModel "github.com/Ho-Minh/InitiaRe-website/internal/models"
	"github.com/Ho-Minh/InitiaRe-website/pkg/utils"

	"github.com/labstack/gommon/log"
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
		log.Errorf("usecase.repo.GetListPaging: %v", err)
		return nil, utils.NewError(constants.STATUS_CODE_INTERNAL_SERVER, "Error when get list todo")
	}
	count, err := u.repo.Count(ctx, queries)
	if err != nil {
		log.Errorf("usecase.repo.Count: %v", err)
		return nil, utils.NewError(constants.STATUS_CODE_INTERNAL_SERVER, "Error when get list todo")
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
	res, err := u.repo.Create(ctx, obj);
	if err != nil {
		log.Errorf("usecase.repo.Create: %v", err)
		return nil, utils.NewError(constants.STATUS_CODE_INTERNAL_SERVER, "Error when create category")
	}
	return res.Export(), nil
}

func (u *usecase) CreateMany(ctx context.Context, userId int, params []*models.SaveRequest) (int, error) {
	return 0, nil
}

func (u *usecase) Update(ctx context.Context, userId int, params *models.SaveRequest) (*models.Response, error) {
	return nil, nil
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