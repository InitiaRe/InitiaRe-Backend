package init

import (
	"InitiaRe-website/config"
	handler "InitiaRe-website/internal/auth/delivery/http"
	"InitiaRe-website/internal/auth/repository"
	authUc "InitiaRe-website/internal/auth/usecase"
	userInfoInit "InitiaRe-website/internal/user_info/init"
	userInfoRepo "InitiaRe-website/internal/user_info/repository"

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
	userInfoInit *userInfoInit.Init,
) *Init {
	repo := repository.InitRepo(db)
	cacheRepo := repository.NewCacheRepo(cache)
	userInfoRepoO := userInfoRepo.InitRepo(db)
	usecase := authUc.InitUsecase(cfg, repo, cacheRepo, userInfoRepoO, userInfoInit.Usecase)
	handler := handler.InitHandler(cfg, usecase)
	return &Init{
		CacheRepository: cacheRepo,
		Repository:      repo,
		Usecase:         usecase,
		Handler:         handler,
	}
}
