package init

import (
	"github.com/Ho-Minh/InitiaRe-website/config"
	"github.com/Ho-Minh/InitiaRe-website/internal/middleware"
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
	mw *middleware.MiddlewareManager,
) *Init {
	repo := repository.NewRepo(db)
	usecase := usecase.NewUseCase(repo)
	handler := handler.NewHandler(cfg, usecase, mw)
	return &Init{
		Repository: repo,
		Usecase:    usecase,
		Handler:    handler,
	}
}
