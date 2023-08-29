package init

import (
	"github.com/Ho-Minh/InitiaRe-website/config"
	"github.com/Ho-Minh/InitiaRe-website/internal/user_info/repository"
	"github.com/Ho-Minh/InitiaRe-website/internal/user_info/usecase"

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
