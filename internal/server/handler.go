package server

import (
	_ "github.com/Ho-Minh/InitiaRe-website/docs"
	initArticle "github.com/Ho-Minh/InitiaRe-website/internal/article/init"
	initAuth "github.com/Ho-Minh/InitiaRe-website/internal/auth/init"
	initMW "github.com/Ho-Minh/InitiaRe-website/internal/middleware/init"
	initTodo "github.com/Ho-Minh/InitiaRe-website/internal/todo/init"
	initCategory "github.com/Ho-Minh/InitiaRe-website/internal/category/init"
	initUser "github.com/Ho-Minh/InitiaRe-website/internal/user/init"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// Map Server Handlers
func (s *Server) MapHandlers(e *echo.Echo) error {

	// Init Auth
	auth := initAuth.NewInit(s.db, s.cfg, s.redisClient)

	// Init Middlewares
	mw := initMW.NewInit(s.cfg, auth)

	// Init Todo
	todo := initTodo.NewInit(s.db, s.cfg, mw)

	// Init Article
	article := initArticle.NewInit(s.db, s.cfg, mw)
	
	// Init Category
	category := initCategory.NewInit(s.db, s.cfg, mw)

	// Init User
	user := initUser.NewInit(s.db, s.cfg, mw)

	v1 := e.Group("/api/v1")
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	authGroup := v1.Group("/auth")
	todoGroup := v1.Group("/todos")
	articleGroup := v1.Group("/articles")
	categoryGroup := v1.Group("/categories")
	userGroup := v1.Group("/users")

	auth.Handler.MapRoutes(authGroup)
	todo.Handler.MapRoutes(todoGroup)
	article.Handler.MapRoutes(articleGroup)
	category.Handler.MapRoutes(categoryGroup)
	user.Handler.MapRoutes(userGroup)

	if s.cfg.Server.Debug {
		log.SetLevel(log.DEBUG)
	}

	return nil
}
