package init

import (
	"github.com/Ho-Minh/InitiaRe-website/config"
	initMW "github.com/Ho-Minh/InitiaRe-website/internal/middleware/init"
	handler "github.com/Ho-Minh/InitiaRe-website/internal/rating/delivery/http"
	"github.com/Ho-Minh/InitiaRe-website/internal/rating/repository"
	"github.com/Ho-Minh/InitiaRe-website/internal/rating/usecase"

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
