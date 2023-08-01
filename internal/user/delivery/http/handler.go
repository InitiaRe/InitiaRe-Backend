package handler

import (
	"net/http"

	"github.com/Ho-Minh/InitiaRe-website/config"
	"github.com/Ho-Minh/InitiaRe-website/internal/middleware"
	"github.com/Ho-Minh/InitiaRe-website/pkg/httpResponse"
	_ "github.com/Ho-Minh/InitiaRe-website/internal/auth/models"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	cfg *config.Config
	mw  middleware.IMiddlewareManager
}

func NewHandler(cfg *config.Config, mw middleware.IMiddlewareManager) IHandler {
	return Handler{
		cfg: cfg,
		mw:  mw,
	}
}

// Map routes
func (h Handler) MapRoutes(group *echo.Group) {
	group.GET("/me", h.GetMe(), h.mw.AuthJWTMiddleware())
}

// GetMe godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Get user info
//	@Description	Get user info by token
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	models.Response
//	@Router			/users/me [get]
func (h Handler) GetMe() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user")
		return c.JSON(http.StatusOK, httpResponse.NewRestResponse(http.StatusOK, "Success", user))
	}
}
