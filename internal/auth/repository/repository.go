package repository

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// Auth repository
type repo struct {
	db *gorm.DB
}

// Constructor
func NewRepo(db *gorm.DB) IRepository {
	return &repo{
		db: db,
	}
}

// Auth redis repository
type redisRepo struct {
	redisClient *redis.Client
}

// Auth redis repository constructor
func NewRedisRepo(redisClient *redis.Client) IRedisRepository {
	return &redisRepo{
		redisClient: redisClient,
	}
}
