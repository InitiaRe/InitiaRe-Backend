package repository

import (
	"context"
	"encoding/json"
	"time"

	"github.com/Ho-Minh/InitiaRe-website/internal/models"
)

// Get user by id
func (a *authRedisRepo) GetByID(ctx context.Context, key string) (*models.User, error) {

	userBytes, err := a.redisClient.Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}
	user := &models.User{}
	if err = json.Unmarshal(userBytes, user); err != nil {
		return nil, err
	}
	return user, nil
}

// Cache user with duration in seconds
func (a *authRedisRepo) SetUser(ctx context.Context, key string, seconds int, user *models.User) error {

	userBytes, err := json.Marshal(user)
	if err != nil {
		return err
	}
	if err = a.redisClient.Set(ctx, key, userBytes, (time.Second * time.Duration(seconds))).Err(); err != nil {
		return err
	}
	return nil
}
