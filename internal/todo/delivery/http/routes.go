package http

import (
	"github.com/labstack/echo/v4"

	"github.com/Ho-Minh/InitiaRe-website/config"
	"github.com/Ho-Minh/InitiaRe-website/internal/auth"
	"github.com/Ho-Minh/InitiaRe-website/internal/middleware"
	"github.com/Ho-Minh/InitiaRe-website/internal/todo"
)

// Map todo routes
func MapTodoRoutes(todoGroup *echo.Group, h todo.Handlers, cfg *config.Config, authUC auth.UseCase, mw *middleware.MiddlewareManager) {
	todoGroup.POST("", h.Create(), mw.AuthJWTMiddleware(cfg, authUC))
	// newsGroup.PUT("/:news_id", h.Update(), mw.AuthSessionMiddleware, mw.CSRF)
	// newsGroup.DELETE("/:news_id", h.Delete(), mw.AuthSessionMiddleware, mw.CSRF)
	// newsGroup.GET("/:news_id", h.GetByID())
	// newsGroup.GET("/search", h.SearchByTitle())
	// newsGroup.GET("", h.GetNews())
}
