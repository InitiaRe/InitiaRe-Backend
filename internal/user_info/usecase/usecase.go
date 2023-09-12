package usecase

import (
	"context"

	"github.com/Ho-Minh/InitiaRe-website/constant"
	"github.com/Ho-Minh/InitiaRe-website/internal/user_info/models"
	"github.com/Ho-Minh/InitiaRe-website/internal/user_info/repository"
	"github.com/Ho-Minh/InitiaRe-website/pkg/utils"
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


func (u *usecase) Enable(ctx context.Context, userId int) (int, error) {
	foundUser, err := u.repo.GetOne(ctx, (&models.RequestList{UserId: userId}).ToMap())
	if err != nil {
		log.Error().Err(err).Str("prefix", "User Info").Str("service", "usecase.repo.GetOne").Send()
		return 0, utils.NewError(constant.STATUS_CODE_BAD_REQUEST, constant.STATUS_MESSAGE_INTERNAL_SERVER_ERROR)
	}

	if foundUser == nil {
		log.Error().Str("prefix", "User Info").Msgf("User not found with userId: %v", userId)
		return 0, utils.NewError(constant.STATUS_CODE_BAD_REQUEST, constant.STATUS_MESSAGE_USER_NOT_FOUND)
	}
	foundUser.Status = constant.USER_STATUS_ACTIVE
	res, err := u.repo.Update(ctx, foundUser);
	if err != nil {
		log.Error().Str("prefix", "User Info").Msgf("Cannot update status with userId: &v", userId)
		return 0, utils.NewError(constant.STATUS_CODE_BAD_REQUEST, constant.STATUS_MESSAGE_INTERNAL_SERVER_ERROR)
	}

	if res == nil {
		log.Error().Str("prefix", "User Info").Msgf("Cannot update status with userId: &v", userId)
		return 0, utils.NewError(constant.STATUS_CODE_BAD_REQUEST, constant.STATUS_MESSAGE_INTERNAL_SERVER_ERROR)
	}

	return 1, nil
}

func (u *usecase) Disable(ctx context.Context, userId int) (int, error) {
	foundUser, err := u.repo.GetOne(ctx, (&models.RequestList{UserId: userId}).ToMap())
	if err != nil {
		log.Error().Err(err).Str("prefix", "User Info").Str("service", "usecase.repo.GetOne").Send()
		return 0, utils.NewError(constant.STATUS_CODE_BAD_REQUEST, constant.STATUS_MESSAGE_INTERNAL_SERVER_ERROR)
	}

	if foundUser == nil {
		log.Error().Str("prefix", "User Info").Msgf("User not found with userId: %v", userId)
		return 0, utils.NewError(constant.STATUS_CODE_BAD_REQUEST, constant.STATUS_MESSAGE_USER_NOT_FOUND)
	}
	foundUser.Status = constant.USER_STATUS_INACTIVE
	res, err := u.repo.Update(ctx, foundUser);
	if err != nil {
		log.Error().Str("prefix", "User Info").Msgf("Cannot update status with userId: &v", userId)
		return 0, utils.NewError(constant.STATUS_CODE_BAD_REQUEST, constant.STATUS_MESSAGE_INTERNAL_SERVER_ERROR)
	}

	if res == nil {
		log.Error().Str("prefix", "User Info").Msgf("Cannot update status with userId: &v", userId)
		return 0, utils.NewError(constant.STATUS_CODE_BAD_REQUEST, constant.STATUS_MESSAGE_INTERNAL_SERVER_ERROR)
	}

	return 1, nil
}
