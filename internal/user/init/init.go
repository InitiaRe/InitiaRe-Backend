package init

import (
	"github.com/Ho-Minh/InitiaRe-website/config"
	"github.com/Ho-Minh/InitiaRe-website/internal/auth/repository"
	initMW "github.com/Ho-Minh/InitiaRe-website/internal/middleware/init"
	handler "github.com/Ho-Minh/InitiaRe-website/internal/user/delivery/http"
	"github.com/Ho-Minh/InitiaRe-website/internal/user/usecase"
	userInfoRepo "github.com/Ho-Minh/InitiaRe-website/internal/user_info/repository"
	"gorm.io/gorm"
)

type Init struct {
	Repository repository.IRepository
	RepositoryUserInfo userInfoRepo.IRepository
	Usecase    usecase.IUseCase
	Handler handler.IHandler
}

func NewInit(
	db *gorm.DB,
	cfg *config.Config,
	mw *initMW.Init,
) *Init {
	repoUserInfo := userInfoRepo.InitRepo(db);
	usecase := usecase.InitUsecase(repoUserInfo);
	handler := handler.InitHandler(cfg, usecase,mw.MiddlewareManager)
	return &Init{
		RepositoryUserInfo: repoUserInfo,
		Usecase:    usecase,
		Handler: handler,
	}
}
