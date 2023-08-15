package init

import (
	"github.com/Ho-Minh/InitiaRe-website/config"
	handler "github.com/Ho-Minh/InitiaRe-website/internal/auth/delivery/http"
	"github.com/Ho-Minh/InitiaRe-website/internal/auth/repository"
	"github.com/Ho-Minh/InitiaRe-website/internal/auth/usecase"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Init struct {
	CacheRepository repository.ICacheRepository
	Repository      repository.IRepository
	Usecase         usecase.IUseCase
	Handler         handler.IHandler
}

func NewInit(
	db *gorm.DB,
	cfg *config.Config,
	cache *redis.Client,
) *Init {
	repo := repository.InitRepo(db)
	cacheRepo := repository.NewCacheRepo(cache)
	usecase := usecase.InitUsecase(cfg, repo, cacheRepo)
	handler := handler.InitHandler(cfg, usecase)
	return &Init{
		CacheRepository: cacheRepo,
		Repository:      repo,
		Usecase:         usecase,
		Handler:         handler,
	}
}
