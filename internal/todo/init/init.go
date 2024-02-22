package init

import (
	"InitiaRe-website/config"
	initMW "InitiaRe-website/internal/middleware/init"
	handler "InitiaRe-website/internal/todo/delivery/http"
	"InitiaRe-website/internal/todo/repository"
	"InitiaRe-website/internal/todo/usecase"

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
	usecase := usecase.InitUsecase(repo)
	handler := handler.InitHandler(cfg, usecase, mw.MiddlewareManager)
	return &Init{
		Repository: repo,
		Usecase:    usecase,
		Handler:    handler,
	}
}
