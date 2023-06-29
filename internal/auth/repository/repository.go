package repository

import (
	"github.com/Ho-Minh/InitiaRe-website/internal/auth"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// Auth repository
type authRepo struct {
	db *gorm.DB
}

// Constructor
func NewAuthRepository(db *gorm.DB) auth.AuthRepository {
	return &authRepo{
		db: db,
	}
}

// Auth redis repository
type authRedisRepo struct {
	redisClient *redis.Client
}

// Auth redis repository constructor
func NewAuthRedisRepo(redisClient *redis.Client) auth.RedisRepository {
	return &authRedisRepo{redisClient: redisClient}
}
