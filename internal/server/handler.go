package server

import (
	_ "github.com/Ho-Minh/InitiaRe-website/docs"
	initAuth "github.com/Ho-Minh/InitiaRe-website/internal/auth/init"
	apiMiddlewares "github.com/Ho-Minh/InitiaRe-website/internal/middleware"
	initTodo "github.com/Ho-Minh/InitiaRe-website/internal/todo/init"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// Map Server Handlers
func (s *Server) MapHandlers(e *echo.Echo) error {

	// Init Auth
	auth := initAuth.NewInit(s.db, s.cfg, s.redisClient)

	// Init middlewares
	mw := apiMiddlewares.NewMiddlewareManager(s.cfg, auth.Repository)

	// Init Todo
	todo := initTodo.NewInit(s.db, s.cfg, mw)

	v1 := e.Group("/api/v1")
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	authGroup := v1.Group("/auth")
	todoGroup := v1.Group("/todo")

	auth.Handler.MapAuthRoutes(authGroup)
	todo.Handler.MapTodoRoutes(todoGroup)

	if s.cfg.Server.Debug {
		log.SetLevel(log.DEBUG)
	}

	return nil
}
