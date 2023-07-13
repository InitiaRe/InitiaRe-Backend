package init

import (
	"github.com/Ho-Minh/InitiaRe-website/config"
	initMW "github.com/Ho-Minh/InitiaRe-website/internal/middleware/init"
	handler "github.com/Ho-Minh/InitiaRe-website/internal/todo/delivery/http"
	"github.com/Ho-Minh/InitiaRe-website/internal/todo/repository"
	"github.com/Ho-Minh/InitiaRe-website/internal/todo/usecase"

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
	repo := repository.NewRepo(db)
	usecase := usecase.NewUseCase(repo)
	handler := handler.NewHandler(cfg, usecase, mw.MiddlewareManager)
	return &Init{
		Repository: repo,
		Usecase:    usecase,
		Handler:    handler,
	}
}
