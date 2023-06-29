package usecase

import (
	"context"

	"github.com/Ho-Minh/InitiaRe-website/internal/constants"
	"github.com/Ho-Minh/InitiaRe-website/internal/models"
	"github.com/Ho-Minh/InitiaRe-website/pkg/utils"

	"github.com/labstack/gommon/log"
)

func (a *authUseCase) SignIn(ctx context.Context, email, password string) (*models.UserWithToken, error) {
	log.SetPrefix("[SignIn]")
	log.Infof("Sign in with user {Email: %s}", email)

	// validation

	// check if user already exists
	foundUser, err := a.authRepo.FindByEmail(ctx, email)
	if err != nil {
		log.Errorf("Error while finding user by email: %s", err)
		return nil, utils.NewError(constants.STATUS_CODE_BAD_REQUEST, constants.STATUS_MESSAGE_INTERNAL_SERVER_ERROR)
	}
	if foundUser == nil {
		log.Errorf("User not found with email: %v", err)
		return nil, utils.NewError(constants.STATUS_CODE_BAD_REQUEST, constants.STATUS_MESSAGE_USER_NOT_FOUND)
	}

	// check if password is correct
	if err = foundUser.ComparePasswords(password); err != nil {
		log.Errorf("Compare password failed: %v", err)
		return nil, utils.NewError(constants.STATUS_CODE_UNAUTHORIZED, constants.STATUS_MESSAGE_INVALID_EMAIL_OR_PASSWORD)
	}
	// end validation

	// generate token
	token, err := utils.GenerateJWTToken(foundUser, a.cfg)
	if err != nil {
		log.Errorf("Cannot generate token: %v", err)
		return nil, utils.NewError(constants.STATUS_CODE_INTERNAL_SERVER, constants.STATUS_MESSAGE_INTERNAL_SERVER_ERROR)
	}

	// save to cache
	if err = a.redisRepo.SetUser(ctx, a.GenerateUserKey(foundUser.UserId.String()), cacheDuration, foundUser); err != nil {
		log.Errorf("authUC.GetByID.SetUserCtx: %v", err)
		return nil, err
	}

	foundUser.SanitizePassword()

	return &models.UserWithToken{
		User:  foundUser,
		Token: token,
	}, nil
}
