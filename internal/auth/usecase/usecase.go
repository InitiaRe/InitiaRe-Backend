package usecase

import (
	"context"
	"time"

	"github.com/Ho-Minh/InitiaRe-website/config"
	"github.com/Ho-Minh/InitiaRe-website/constant"
	authEntity "github.com/Ho-Minh/InitiaRe-website/internal/auth/entity"
	authModel "github.com/Ho-Minh/InitiaRe-website/internal/auth/models"
	"github.com/Ho-Minh/InitiaRe-website/internal/auth/repository"
	userInfoModel "github.com/Ho-Minh/InitiaRe-website/internal/user_info/models"
	userInfoRepo "github.com/Ho-Minh/InitiaRe-website/internal/user_info/repository"
	userInfoUc "github.com/Ho-Minh/InitiaRe-website/internal/user_info/usecase"
	"github.com/Ho-Minh/InitiaRe-website/pkg/utils"

	"github.com/rs/zerolog/log"
)

type usecase struct {
	cfg          *config.Config
	repo         repository.IRepository
	userInfoRepo userInfoRepo.IRepository
	cacheRepo    repository.ICacheRepository
	userInfoUc   userInfoUc.IUseCase
}

func InitUsecase(cfg *config.Config, repo repository.IRepository, cacheRepo repository.ICacheRepository, userInfoRepo userInfoRepo.IRepository, userInfoUc userInfoUc.IUseCase) IUseCase {
	return &usecase{
		cfg:          cfg,
		repo:         repo,
		userInfoRepo: userInfoRepo,
		cacheRepo:    cacheRepo,
		userInfoUc:   userInfoUc,
	}
}

func (u *usecase) Register(ctx context.Context, params *authModel.SaveRequest) (*authModel.Response, error) {
	log.Info().Str("prefix", "Auth").Msgf("Register user with params: {FirstName: %s, LastName: %s, Email: %s}", params.FirstName, params.LastName, params.Email)

	// validation

	// check if user already exists
	foundUser, err := u.repo.GetOne(ctx, (&authModel.RequestList{Email: params.Email}).ToMap())
	if err != nil {
		log.Error().Err(err).Str("prefix", "Auth").Str("service", "usecase.repo.GetOne").Send()
		return nil, utils.NewError(constant.STATUS_CODE_BAD_REQUEST, constant.STATUS_MESSAGE_INTERNAL_SERVER_ERROR)
	}
	if foundUser.Id != 0 {
		log.Error().Str("prefix", "Auth").Msgf("User already exist with email: %v", params.Email)
		return nil, utils.NewError(constant.STATUS_CODE_BAD_REQUEST, constant.STATUS_MESSAGE_EMAIL_ALREADY_EXISTS)
	}

	if params.Gender != "Male" && params.Gender != "Female" {
		log.Error().Str("prefix", "Auth").Msgf("Invalid gender type: %s", params.Gender)
		return nil, utils.NewError(constant.STATUS_CODE_BAD_REQUEST, constant.STATUS_MESSAGE_INVALID_GENDER_TYPE)
	}
	// end validation

	// create new user
	user := &authEntity.User{}
	user.HashPassword()
	user.ParseFromSaveRequest(params)
	resUser, err := u.repo.Create(ctx, user)
	if err != nil {
		log.Error().Err(err).Str("prefix", "Auth").Str("service", "usecase.repo.Create").Send()
		return nil, err
	}
	resUser.SanitizePassword()

	// create new user info
	if _, err := u.userInfoUc.Create(ctx, resUser.Id, &userInfoModel.SaveRequest{
		UserId: resUser.Id,
	}); err != nil {
		return nil, err
	}
	return resUser.Export(), nil
}

func (u *usecase) Login(ctx context.Context, params *authModel.LoginRequest) (*authModel.UserWithToken, error) {
	log.Info().Str("prefix", "Auth").Msgf("Sign in with user {Email: %v}", params.Email)

	// validation

	// check if user already exists
	foundUser, err := u.repo.GetOne(ctx, (&authModel.RequestList{Email: params.Email}).ToMap())
	if err != nil {
		log.Error().Err(err).Str("prefix", "Auth").Str("service", "usecase.repo.GetOne").Send()
		return nil, utils.NewError(constant.STATUS_CODE_BAD_REQUEST, constant.STATUS_MESSAGE_INTERNAL_SERVER_ERROR)
	}
	if foundUser == nil {
		log.Error().Str("prefix", "Auth").Msgf("User not found with email: %v", params.Email)
		return nil, utils.NewError(constant.STATUS_CODE_BAD_REQUEST, constant.STATUS_MESSAGE_USER_NOT_FOUND)
	}

	// check if password is correct
	if err = utils.ComparePassword(foundUser.Password, params.Password); err != nil {
		log.Error().Err(err).Str("prefix", "Auth").Str("service", "utils.ComparePassword").Send()
		return nil, utils.NewError(constant.STATUS_CODE_UNAUTHORIZED, constant.STATUS_MESSAGE_INVALID_EMAIL_OR_PASSWORD)
	}
	// end validation

	foundUserInfo, err := u.userInfoRepo.GetOne(ctx, (&userInfoModel.RequestList{UserId: foundUser.Id}).ToMap())
	if err != nil {
		log.Error().Err(err).Str("prefix", "Auth").Str("service", "usecase.userInfoRepo.GetOne").Send()
		return nil, utils.NewError(constant.STATUS_CODE_BAD_REQUEST, constant.STATUS_MESSAGE_INTERNAL_SERVER_ERROR)
	}
	if foundUserInfo == nil {
		log.Error().Str("prefix", "Auth").Msgf("User not found with userId: %v", foundUser.Id)
		return nil, utils.NewError(constant.STATUS_CODE_BAD_REQUEST, constant.STATUS_MESSAGE_USER_NOT_FOUND)
	}

	if foundUserInfo.Status == constant.USER_STATUS_INACTIVE {
		log.Error().Str("prefix", "Auth").Msgf("User is not activated with userId: %v", foundUser.Id)
		return nil, utils.NewError(constant.STATUS_CODE_FORBIDDEN, constant.STATUS_MESSAGE_USER_INACTIVE)
	}

	// generate token
	token, err := utils.GenerateJWTToken(foundUser.Export(), u.cfg.Auth.Secret, u.cfg.Auth.Expire)
	if err != nil {
		log.Error().Err(err).Str("prefix", "Auth").Str("service", "utils.GenerateJWTToken").Send()
		return nil, utils.NewError(constant.STATUS_CODE_INTERNAL_SERVER, constant.STATUS_MESSAGE_INTERNAL_SERVER_ERROR)
	}

	// save to cache
	if err = u.cacheRepo.SetUser(ctx, utils.GenerateUserKey(foundUser.Id), u.cfg.Auth.Expire, foundUser); err != nil {
		log.Error().Err(err).Str("prefix", "Auth").Str("service", "usecase.redisRepo.SetUser").Send()
		return nil, err
	}

	foundUser.SanitizePassword()
	foundUser.LoginDate = time.Now()
	if _, err := u.repo.Update(ctx, foundUser); err != nil {
		log.Error().Str("prefix", "Auth").Msgf("Cannot update login_date with userId: %v", foundUser.Id)
		return nil, utils.NewError(constant.STATUS_CODE_INTERNAL_SERVER, constant.STATUS_MESSAGE_INTERNAL_SERVER_ERROR)
	}

	return &authModel.UserWithToken{
		User:  foundUser.Export(),
		Token: token,
	}, nil
}

func (u *usecase) GetOne(ctx context.Context, params *authModel.RequestList) (*authModel.Response, error) {
	queries := params.ToMap()
	record, err := u.repo.GetOne(ctx, queries)
	if err != nil {
		log.Error().Err(err).Str("prefix", "Auth").Str("service", "usecase.repo.GetOne").Send()
		return nil, utils.NewError(constant.STATUS_CODE_INTERNAL_SERVER, "Error when get user")
	}
	return record.Export(), nil
}