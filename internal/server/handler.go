package server

import (
	_ "github.com/Ho-Minh/InitiaRe-website/docs"
	authHttp "github.com/Ho-Minh/InitiaRe-website/internal/auth/delivery/http"
	authRepository "github.com/Ho-Minh/InitiaRe-website/internal/auth/repository"
	authUseCase "github.com/Ho-Minh/InitiaRe-website/internal/auth/usecase"
	apiMiddlewares "github.com/Ho-Minh/InitiaRe-website/internal/middleware"
	todoHttp "github.com/Ho-Minh/InitiaRe-website/internal/todo/delivery/http"
	todoRepository "github.com/Ho-Minh/InitiaRe-website/internal/todo/repository"
	todoUseCase "github.com/Ho-Minh/InitiaRe-website/internal/todo/usecase"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// Map Server Handlers
func (s *Server) MapHandlers(e *echo.Echo) error {

	// Init repositories
	authRepo := authRepository.NewAuthRepository(s.db)
	authRedisRepo := authRepository.NewAuthRedisRepo(s.redisClient)
	todoRepo := todoRepository.NewTodoRepository(s.db)

	// Init useCases
	authUC := authUseCase.NewAuthUseCase(s.cfg, authRepo, authRedisRepo)
	todoUC := todoUseCase.NewTodoUseCase(todoRepo)

	// Init handlers
	authHandlers := authHttp.NewAuthHandlers(s.cfg, authUC)
	todoHandlers := todoHttp.NewTodoHandlers(s.cfg, todoUC)

	// Init middlewares
	mw := apiMiddlewares.NewMiddlewareManager(s.cfg, authUC)

	v1 := e.Group("/api/v1")
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	// e.Use(middleware.Logger())

	authGroup := v1.Group("/auth")
	todoGroup := v1.Group("/todo")

	authHttp.MapAuthRoutes(authGroup, authHandlers)
	todoHttp.MapTodoRoutes(todoGroup, todoHandlers, s.cfg, authUC, mw)

	if s.cfg.Server.Debug {
		log.SetLevel(log.DEBUG)
	}

	return nil
}
