package usecase

import (
	"context"

	"github.com/Ho-Minh/InitiaRe-website/internal/constants"
	"github.com/Ho-Minh/InitiaRe-website/internal/models"
	"github.com/Ho-Minh/InitiaRe-website/pkg/utils"

	"github.com/labstack/gommon/log"
)

func (a *authUseCase) SignUp(ctx context.Context, user *models.User) (*models.User, error) {
	log.SetPrefix("[SignUp]")
	log.Infof("Sign up new user with params: {FirstName: %s, LastName: %s, Email: %s}", user.FirstName, user.LastName, user.Email)

	// validation

	// check if user already exists
	foundUser, err := a.authRepo.FindByEmail(ctx, user.Email)
	if err != nil {
		log.Errorf("Error while finding user by email: %s", err)
		return nil, utils.NewError(constants.STATUS_CODE_BAD_REQUEST, constants.STATUS_MESSAGE_INTERNAL_SERVER_ERROR)
	}
	if foundUser != nil {
		log.Errorf("User already exist with email: %v", err)
		return nil, utils.NewError(constants.STATUS_CODE_BAD_REQUEST, constants.STATUS_MESSAGE_EMAIL_ALREADY_EXISTS)
	}

	if user.Gender != "Male" && user.Gender != "Female" {
		log.Errorf("Invalid gender type: %s", user.Gender)
		return nil, utils.NewError(constants.STATUS_CODE_BAD_REQUEST, constants.STATUS_MESSAGE_INVALID_GENDER_TYPE)
	}
	// end validation

	// create new user
	_, err = a.authRepo.SignUp(ctx, user.Parse())
	if err != nil {
		log.Errorf("Error while creating new user: %s", err)
		return nil, err
	}

	return user, nil
}
