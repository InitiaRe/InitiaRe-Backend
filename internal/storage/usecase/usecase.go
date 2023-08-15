package usecase

import (
	"context"

	"github.com/Ho-Minh/InitiaRe-website/config"
	"github.com/Ho-Minh/InitiaRe-website/constant"
	"github.com/Ho-Minh/InitiaRe-website/internal/storage/entity"
	"github.com/Ho-Minh/InitiaRe-website/internal/storage/models"
	"github.com/Ho-Minh/InitiaRe-website/internal/storage/repository"
	"github.com/Ho-Minh/InitiaRe-website/pkg/utils"

	"github.com/rs/zerolog/log"
)

type usecase struct {
	cfg     *config.Config
	repo    repository.IRepository
	ctnRepo repository.IContainerRepository
}

func InitUsecase(cfg *config.Config, repo repository.IRepository, ctnRepo repository.IContainerRepository) IUseCase {
	return &usecase{
		cfg:     cfg,
		repo:    repo,
		ctnRepo: ctnRepo,
	}
}
func (u *usecase) UploadMedia(ctx context.Context, userId int, params *models.UploadRequest) (*models.Response, error) {
	log.Info().Str("prefix", "Storage").Msgf("Upload file by user [%v] with params: %+v", userId, params)
	if err := u.validateBeforeUpload(ctx, params.Obj); err != nil {
		return nil, err
	}

	url, err := u.ctnRepo.Upload(ctx, params)
	if err != nil {
		log.Error().Err(err).Str("prefix", "Storage").Str("service", "usecase.ctnRepo.Upload").Send()
		return nil, utils.NewError(constant.STATUS_CODE_INTERNAL_SERVER, "Error when upload file")
	}

	obj := &entity.Storage{}
	obj.ParseForCreate(&models.SaveRequest{
		DownloadUrl: url,
		Type:        constant.STORAGE_TYPE_MEDIA,
	}, userId)

	res, err := u.repo.Create(ctx, obj)
	if err != nil {
		log.Error().Err(err).Str("prefix", "Storage").Str("service", "usecase.repo.Create").Send()
		return nil, utils.NewError(constant.STATUS_CODE_INTERNAL_SERVER, "Error when create storage history")
	}
	return res.Export(), nil
}

func (u *usecase) validateBeforeUpload(ctx context.Context, obj []byte) error {
	// record, err := u.repo.GetById(ctx, id)
	// if err != nil {
	// 	log.Error().Err(err).Str("prefix", "Todo").Str("service", "usecase.repo.GetById")
	// 	return utils.NewError(constant.STATUS_CODE_INTERNAL_SERVER, "Error when get todo")
	// }
	// if record.Id == 0 {
	// 	return utils.NewError(constant.STATUS_CODE_NOT_FOUND, "Todo not found")
	// }
	return nil
}
