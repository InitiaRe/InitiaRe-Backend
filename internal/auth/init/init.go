package init

import (
	"github.com/Ho-Minh/InitiaRe-website/config"
	handler "github.com/Ho-Minh/InitiaRe-website/internal/auth/delivery/http"
	"github.com/Ho-Minh/InitiaRe-website/internal/auth/repository"
	authUc "github.com/Ho-Minh/InitiaRe-website/internal/auth/usecase"
	userInfoUc "github.com/Ho-Minh/InitiaRe-website/internal/user_info/usecase"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Init struct {
	CacheRepository repository.ICacheRepository
	Repository      repository.IRepository
	Usecase         authUc.IUseCase
	Handler         handler.IHandler
}

func NewInit(
	db *gorm.DB,
	cfg *config.Config,
	cache *redis.Client,
	userInfoUc userInfoUc.IUseCase,
) *Init {
	repo := repository.InitRepo(db)
	cacheRepo := repository.NewCacheRepo(cache)
	usecase := authUc.InitUsecase(cfg, repo, cacheRepo, userInfoUc)
	handler := handler.InitHandler(cfg, usecase)
	return &Init{
		CacheRepository: cacheRepo,
		Repository:      repo,
		Usecase:         usecase,
		Handler:         handler,
	}
}
