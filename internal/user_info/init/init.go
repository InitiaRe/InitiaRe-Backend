package init

import (
	"InitiaRe-website/config"
	handler "InitiaRe-website/internal/user_info/delivery/http"
	"InitiaRe-website/internal/user_info/repository"
	"InitiaRe-website/internal/user_info/usecase"

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
		Handler:    handler,
	}
}
