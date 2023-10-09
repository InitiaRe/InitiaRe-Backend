package handler

import (
	"net/http"

	"github.com/Ho-Minh/InitiaRe-website/config"
	"github.com/Ho-Minh/InitiaRe-website/constant"
	"github.com/Ho-Minh/InitiaRe-website/internal/auth/models"
	"github.com/Ho-Minh/InitiaRe-website/internal/middleware"
	"github.com/Ho-Minh/InitiaRe-website/internal/user/usecase"
	"github.com/Ho-Minh/InitiaRe-website/pkg/httpResponse"
	"github.com/Ho-Minh/InitiaRe-website/pkg/utils"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	cfg     *config.Config
	mw      middleware.IMiddlewareManager
	usecase usecase.IUseCase
}

func InitHandler(cfg *config.Config, mw middleware.IMiddlewareManager, usecase usecase.IUseCase) IHandler {
	return Handler{
		cfg:     cfg,
		mw:      mw,
		usecase: usecase,
	}
}

// Map routes
func (h Handler) MapRoutes(group *echo.Group) {
	group.GET("/me", h.GetMe(), h.mw.AuthJWTMiddleware())
	group.POST("/enable", h.Enable(), h.mw.AuthJWTMiddleware())
	group.PUT("/disable", h.Disable(), h.mw.AuthJWTMiddleware())
}

// GetMe godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Get user info
//	@Description	Get user info by token
//	@Tags			User
//	@Produce		json
//	@Success		200	{object}	models.Response
//	@Router			/users/me [get]
func (h Handler) GetMe() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user")
		return c.JSON(http.StatusOK, httpResponse.NewRestResponse(http.StatusOK, "Success", user))
	}
}

func (h Handler) Enable() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := utils.GetRequestCtx(c)
		user := c.Get("user").(*models.Response)
		if err := h.usecase.Enable(ctx, user.Id); err != nil {
			return c.JSON(http.StatusOK, httpResponse.ParseError(err))
		}
		return c.JSON(http.StatusOK, httpResponse.NewRestResponse(http.StatusOK, constant.STATUS_MESSAGE_OK, 1))
	}
}

func (h Handler) Disable() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := utils.GetRequestCtx(c)
		user := c.Get("user").(*models.Response)
		if err := h.usecase.Disable(ctx, user.Id); err != nil {
			return c.JSON(http.StatusOK, httpResponse.ParseError(err))
		}
		return c.JSON(http.StatusOK, httpResponse.NewRestResponse(http.StatusOK, constant.STATUS_MESSAGE_OK, 1))
	}
}
