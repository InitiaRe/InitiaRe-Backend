package init

import (
	"github.com/Ho-Minh/InitiaRe-website/config"
	initAuth "github.com/Ho-Minh/InitiaRe-website/internal/auth/init"
	initMW "github.com/Ho-Minh/InitiaRe-website/internal/middleware/init"
	handler "github.com/Ho-Minh/InitiaRe-website/internal/user/delivery/http"
	"github.com/Ho-Minh/InitiaRe-website/internal/user/usecase"
	initUserInfo "github.com/Ho-Minh/InitiaRe-website/internal/user_info/init"
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
