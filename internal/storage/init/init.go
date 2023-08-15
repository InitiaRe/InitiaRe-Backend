package init

import (
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Ho-Minh/InitiaRe-website/config"
	initMW "github.com/Ho-Minh/InitiaRe-website/internal/middleware/init"
	handler "github.com/Ho-Minh/InitiaRe-website/internal/storage/delivery/http"
	"github.com/Ho-Minh/InitiaRe-website/internal/storage/repository"
	"github.com/Ho-Minh/InitiaRe-website/internal/storage/usecase"

	"gorm.io/gorm"
)

type Init struct {
	CtnRepository repository.IContainerRepository
	Repository    repository.IRepository
	Usecase       usecase.IUseCase
	Handler       handler.IHandler
}

func NewInit(
	db *gorm.DB,
	cfg *config.Config,
	mw *initMW.Init,
	ctn *azblob.Client,
) *Init {
	repo := repository.InitRepo(db)
	ctnRepo := repository.InitContainerRepo(cfg, ctn)
	usecase := usecase.InitUsecase(cfg, repo, ctnRepo)
	handler := handler.InitHandler(cfg, usecase, mw.MiddlewareManager)
	return &Init{
		Repository: repo,
		Usecase:    usecase,
		Handler:    handler,
	}
}
