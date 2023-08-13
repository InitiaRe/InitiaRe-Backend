package repository

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) IRepository {
	return &repo{
		db: db,
	}
}

type cacheRepo struct {
	cache *redis.Client
}

func NewCacheRepo(cache *redis.Client) ICacheRepository {
	return &cacheRepo{
		cache: cache,
	}
}
