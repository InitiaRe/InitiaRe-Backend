package handler

import (
	"net/http"

	"github.com/Ho-Minh/InitiaRe-website/config"
	"github.com/Ho-Minh/InitiaRe-website/constant"
	userModel "github.com/Ho-Minh/InitiaRe-website/internal/auth/models"
	"github.com/Ho-Minh/InitiaRe-website/internal/middleware"
	"github.com/Ho-Minh/InitiaRe-website/internal/user_info/usecase"
	"github.com/Ho-Minh/InitiaRe-website/pkg/httpResponse"
	"github.com/Ho-Minh/InitiaRe-website/pkg/utils"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	cfg     *config.Config
	usecase usecase.IUseCase
	mw      middleware.IMiddlewareManager
}

func InitHandler(cfg *config.Config, usecase usecase.IUseCase, mw middleware.IMiddlewareManager) IHandler {
	return Handler{
		cfg:     cfg,
		usecase: usecase,
		mw:      mw,
	}
}

// Map routes
func (h Handler) MapRoutes(group *echo.Group) {
	group.POST("/enable", h.Enable(), h.mw.AuthJWTMiddleware())
	group.PUT("/disable", h.Disable(), h.mw.AuthJWTMiddleware())
}

func (h Handler) Enable() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := utils.GetRequestCtx(c)
		user := c.Get("user").(*userModel.Response)
		if err := h.usecase.Enable(ctx, user.Id); err != nil {
			return c.JSON(http.StatusOK, httpResponse.ParseError(err))
		}
		return c.JSON(http.StatusOK, httpResponse.NewRestResponse(http.StatusOK, constant.STATUS_MESSAGE_OK, nil))
	}
}

func (h Handler) Disable() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := utils.GetRequestCtx(c)
		user := c.Get("user").(*userModel.Response)
		if err := h.usecase.Disable(ctx, user.Id); err != nil {
			return c.JSON(http.StatusOK, httpResponse.ParseError(err))
		}
		return c.JSON(http.StatusOK, httpResponse.NewRestResponse(http.StatusOK, constant.STATUS_MESSAGE_OK, nil))
	}
}
