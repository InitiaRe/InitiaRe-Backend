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
	RedisRepository repository.IRedisRepository
	Repository      repository.IRepository
	Usecase         usecase.IUseCase
	Handler         handler.IHandler
}

func NewInit(
	db *gorm.DB,
	cfg *config.Config,
	redisClient *redis.Client,
) *Init {
	repo := repository.NewRepo(db)
	redisRepo := repository.NewRedisRepo(redisClient)
	usecase := usecase.NewUseCase(cfg, repo, redisRepo)
	handler := handler.NewHandler(cfg, usecase)
	return &Init{
		RedisRepository: redisRepo,
		Repository:      repo,
		Usecase:         usecase,
		Handler:         handler,
	}
}
