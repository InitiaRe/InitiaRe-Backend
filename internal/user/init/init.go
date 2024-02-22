package init

import (
	"InitiaRe-website/config"
	initAuth "InitiaRe-website/internal/auth/init"
	initMW "InitiaRe-website/internal/middleware/init"
	handler "InitiaRe-website/internal/user/delivery/http"
	"InitiaRe-website/internal/user/usecase"
	initUserInfo "InitiaRe-website/internal/user_info/init"
	"InitiaRe-website/internal/user_info/repository"

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
	initMW *initMW.Init,
	initAuth *initAuth.Init,
	initUserInfo *initUserInfo.Init,
) *Init {
	usecase := usecase.InitUsecase(initAuth.Usecase, initUserInfo.Usecase)
	handler := handler.InitHandler(cfg, initMW.MiddlewareManager, usecase)
	return &Init{
		Usecase: usecase,
		Handler: handler,
	}
}
