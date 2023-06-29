package usecase

import (
	"context"

	"github.com/Ho-Minh/InitiaRe-website/internal/models"

	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
)

func (a *authUseCase) GetByID(ctx context.Context, userID uuid.UUID) (*models.User, error) {

	cachedUser, err := a.redisRepo.GetByID(ctx, a.GenerateUserKey(userID.String()))
	if err != nil {
		log.Errorf("authUC.GetByID.GetByIDCtx: %v", err)
		return nil, err
	}
	if cachedUser != nil {
		return cachedUser, nil
	}

	user, err := a.authRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	if err = a.redisRepo.SetUser(ctx, a.GenerateUserKey(userID.String()), cacheDuration, user); err != nil {
		log.Errorf("authUC.GetByID.SetUserCtx: %v", err)
		return nil, err
	}

	user.SanitizePassword()

	return user, nil
}
