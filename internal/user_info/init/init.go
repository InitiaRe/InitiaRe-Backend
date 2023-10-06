package init

import (
	"github.com/Ho-Minh/InitiaRe-website/config"
	handler "github.com/Ho-Minh/InitiaRe-website/internal/user_info/delivery/http"
	"github.com/Ho-Minh/InitiaRe-website/internal/user_info/repository"
	"github.com/Ho-Minh/InitiaRe-website/internal/user_info/usecase"

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
) *Init {
	repo := repository.InitRepo(db)
	usecase := usecase.InitUsecase(repo)
	handler := handler.InitHandler(cfg, usecase)
	return &Init{
		Repository: repo,
		Usecase:    usecase,
		Handler: handler,
	}
}
