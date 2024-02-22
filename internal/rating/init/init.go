package init

import (
	"InitiaRe-website/config"
	initMW "InitiaRe-website/internal/middleware/init"
	handler "InitiaRe-website/internal/rating/delivery/http"
	"InitiaRe-website/internal/rating/repository"
	"InitiaRe-website/internal/rating/usecase"

	"gorm.io/gorm"
)

type Init struct {
	Repository repository.IRepository
	Usecase    usecase.IUseCase
	Handler    handler.IHandler
}

func NewInit(
	db *gorm.DB,
	cfg *config.Config,
	mw *initMW.Init,
) *Init {
	repo := repository.InitRepo(db)
	usecase := usecase.InitUsecase(cfg, repo)
	handler := handler.InitHandler(cfg, usecase, mw.MiddlewareManager)
	return &Init{
		Repository: repo,
		Usecase:    usecase,
		Handler:    handler,
	}
}
