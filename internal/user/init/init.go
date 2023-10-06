package init

import (
	"github.com/Ho-Minh/InitiaRe-website/config"
	initMW "github.com/Ho-Minh/InitiaRe-website/internal/middleware/init"
	handler "github.com/Ho-Minh/InitiaRe-website/internal/user/delivery/http"
	"github.com/Ho-Minh/InitiaRe-website/internal/user/usecase"
	"github.com/Ho-Minh/InitiaRe-website/internal/user_info/repository"
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
	handler := handler.InitHandler(cfg, mw.MiddlewareManager, usecase)
	return &Init{
		Usecase: usecase,
		Handler: handler,
	}
}
