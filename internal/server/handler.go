package server

import (
	_ "InitiaRe-website/docs"
	initArticle "InitiaRe-website/internal/article/init"
	initArticleCategory "InitiaRe-website/internal/article_category/init"
	initAuth "InitiaRe-website/internal/auth/init"
	initCategory "InitiaRe-website/internal/category/init"
	initMW "InitiaRe-website/internal/middleware/init"
	initRating "InitiaRe-website/internal/rating/init"
	initStorage "InitiaRe-website/internal/storage/init"
	initTodo "InitiaRe-website/internal/todo/init"
	initUser "InitiaRe-website/internal/user/init"
	initUserInfo "InitiaRe-website/internal/user_info/init"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// Map Server Handlers
func (s *Server) MapHandlers(e *echo.Echo) error {

	// Init UserInfo
	userInfo := initUserInfo.NewInit(s.db, s.cfg)

	// Init Auth
	auth := initAuth.NewInit(s.db, s.cfg, s.redisClient, userInfo)

	// Init Middlewares
	mw := initMW.NewInit(s.cfg, auth)

	// Init Todo
	todo := initTodo.NewInit(s.db, s.cfg, mw)

	// Init Rating
	rating := initRating.NewInit(s.db, s.cfg, mw)

	// Init Category
	category := initCategory.NewInit(s.db, s.cfg, mw)

	// Init Article Category
	articleCategory := initArticleCategory.NewInit(s.db, s.cfg)

	// Init Article
	article := initArticle.NewInit(s.db, s.cfg, mw, rating, category, articleCategory)

	// Init User
	user := initUser.NewInit(s.db, s.cfg, mw, auth, userInfo)

	// Init Storage
	storage := initStorage.NewInit(s.db, s.cfg, mw, s.ctn)

	// Enable cors
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	// Group api paths
	v1 := e.Group("/api/v1")
	authGroup := v1.Group("/auth")
	todoGroup := v1.Group("/todos")
	articleGroup := v1.Group("/articles")
	categoryGroup := v1.Group("/categories")
	userGroup := v1.Group("/users")
	storageGroup := v1.Group("/storage")
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	auth.Handler.MapRoutes(authGroup)
	todo.Handler.MapRoutes(todoGroup)
	article.Handler.MapRoutes(articleGroup)
	category.Handler.MapRoutes(categoryGroup)
	user.Handler.MapRoutes(userGroup)
	storage.Handler.MapRoutes(storageGroup)

	return nil
}
