package init

import (
	"InitiaRe-website/config"
	"InitiaRe-website/internal/article_category/repository"
	"InitiaRe-website/internal/article_category/usecase"

	"gorm.io/gorm"
)

type Init struct {
	Repository repository.IRepository
	Usecase    usecase.IUseCase
}

func NewInit(
	db *gorm.DB,
	cfg *config.Config,
) *Init {
	repo := repository.InitRepo(db)
	usecase := usecase.InitUsecase(repo)
	return &Init{
		Repository: repo,
		Usecase:    usecase,
	}
}
