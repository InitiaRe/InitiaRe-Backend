package repository

import (
	"context"
	"time"

	"github.com/Ho-Minh/InitiaRe-website/internal/models"

	"github.com/google/uuid"
)

func (a *authRepo) SignUp(ctx context.Context, user *models.User) (int64, error) {

	if user.Birthday.IsZero() {
		user.Birthday = time.Now().Truncate(24 * time.Hour)
	}

	if err := a.db.Create(&user).Error; err != nil {
		return 0, err
	}
	return 1, nil
}

func (a *authRepo) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	user := &models.User{}
	if err := a.db.Where("email = ?", email).Find(&user).Limit(1).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (a *authRepo) GetByID(ctx context.Context, userID uuid.UUID) (*models.User, error) {
	var user models.User
	if result := a.db.Where("id = ?", userID).First(&user); result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
