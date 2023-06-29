package auth

import (
	"context"

	"github.com/Ho-Minh/InitiaRe-website/internal/models"

	"github.com/google/uuid"
)

// Auth repository interface
type AuthRepository interface {
	SignUp(ctx context.Context, user *models.User) (int64, error)
	// Update(ctx context.Context, user *models.User) (*models.User, error)
	// Delete(ctx context.Context, userID uuid.UUID) error
	GetByID(ctx context.Context, userID uuid.UUID) (*models.User, error)
	// FindByName(ctx context.Context, name string, query *utils.PaginationQuery) (*models.UsersList, error)
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	// GetUsers(ctx context.Context, pq *utils.PaginationQuery) (*models.UsersList, error)
}

// Auth Redis repository interface
type RedisRepository interface {
	GetByID(ctx context.Context, key string) (*models.User, error)
	SetUser(ctx context.Context, key string, seconds int, user *models.User) error
	// DeleteUser(ctx context.Context, key string) error
}
