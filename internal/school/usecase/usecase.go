package usecase

import (
	"InitiaRe-website/constant"
	commonModel "InitiaRe-website/internal/models"
	"InitiaRe-website/internal/school/entity"
	"InitiaRe-website/internal/school/models"
	"InitiaRe-website/internal/school/repository"
	"InitiaRe-website/pkg/utils"
	"context"

	"github.com/rs/zerolog/log"
)

type usecase struct {
	repo repository.IRepository
}

func InitUsecase(
	repo repository.IRepository,
) IUseCase {
	return &usecase{
		repo: repo,
	}
}

func (u *usecase) GetListPaging(ctx context.Context, params *models.RequestList) (*models.ListPaging, error) {
	queries := params.ToMap()
	records, err := u.repo.GetListPaging(ctx, queries)
	if err != nil {
		log.Error().Err(err).Str("service", "usecase.repo.GetListPaging").Send()
		return nil, utils.NewError(constant.STATUS_CODE_INTERNAL_SERVER, "Error when get list school")
	}
	count, err := u.repo.Count(ctx, queries)
	if err != nil {
		log.Error().Err(err).Str("service", "usecase.repo.Count").Send()
		return nil, utils.NewError(constant.STATUS_CODE_INTERNAL_SERVER, "Error when get list school")
	}

	return &models.ListPaging{
		ListPaging: commonModel.ListPaging{
			Page:  params.Page,
			Size:  params.Size,
			Total: count,
		},
		Records: (&entity.School{}).ExportList(records),
	}, nil
}
