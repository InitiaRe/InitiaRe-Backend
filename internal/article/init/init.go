package init

import (
	"InitiaRe-website/config"
	handler "InitiaRe-website/internal/article/delivery/http"
	"InitiaRe-website/internal/article/repository"
	"InitiaRe-website/internal/article/usecase"
	initArticleCategory "InitiaRe-website/internal/article_category/init"
	initCategory "InitiaRe-website/internal/category/init"
	initMW "InitiaRe-website/internal/middleware/init"
	initRating "InitiaRe-website/internal/rating/init"

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
