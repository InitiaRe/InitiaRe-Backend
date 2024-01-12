package init

import (
	"github.com/Ho-Minh/InitiaRe-website/config"
	handler "github.com/Ho-Minh/InitiaRe-website/internal/article/delivery/http"
	"github.com/Ho-Minh/InitiaRe-website/internal/article/repository"
	"github.com/Ho-Minh/InitiaRe-website/internal/article/usecase"
	initArticleCategory "github.com/Ho-Minh/InitiaRe-website/internal/article_category/init"
	initCategory "github.com/Ho-Minh/InitiaRe-website/internal/category/init"
	initMW "github.com/Ho-Minh/InitiaRe-website/internal/middleware/init"
	initRating "github.com/Ho-Minh/InitiaRe-website/internal/rating/init"

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
	initRating *initRating.Init,
	initCategory *initCategory.Init,
	initArticleCategory *initArticleCategory.Init,
) *Init {
	repo := repository.InitRepo(db)
	usecase := usecase.InitUsecase(repo, initRating.Usecase, initCategory.Usecase, initArticleCategory.Usecase)
	handler := handler.InitHandler(cfg, usecase, mw.MiddlewareManager)
	return &Init{
		Repository: repo,
		Usecase:    usecase,
		Handler:    handler,
	}
}
