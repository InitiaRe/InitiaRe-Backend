package handler

import (
	"InitiaRe-website/config"
	"InitiaRe-website/internal/user_info/usecase"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	cfg     *config.Config
	usecase usecase.IUseCase
}

func InitHandler(cfg *config.Config, usecase usecase.IUseCase) IHandler {
	return Handler{
		cfg:     cfg,
		usecase: usecase,
	}
}

// Map routes
func (h Handler) MapRoutes(group *echo.Group) {
}
