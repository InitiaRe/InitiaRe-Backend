package handler

import (
	"net/http"

	"github.com/Ho-Minh/InitiaRe-website/config"
	"github.com/Ho-Minh/InitiaRe-website/constant"
	authModel "github.com/Ho-Minh/InitiaRe-website/internal/auth/models"
	"github.com/Ho-Minh/InitiaRe-website/internal/middleware"
	userModel "github.com/Ho-Minh/InitiaRe-website/internal/user/models"
	"github.com/Ho-Minh/InitiaRe-website/internal/user/usecase"
	"github.com/Ho-Minh/InitiaRe-website/pkg/httpResponse"
	"github.com/Ho-Minh/InitiaRe-website/pkg/utils"
	"github.com/rs/zerolog/log"

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
	group.POST("/promote/admin", h.PromoteAdmin(), h.mw.AuthJWTMiddleware())
}

// GetMe godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Get user info
//	@Description	Get user info by token
//	@Tags			User
//	@Produce		json
//	@Success		200	{object}	authModel.Response
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
		user := c.Get("user").(*authModel.Response)
		if err := h.usecase.Enable(ctx, user.Id); err != nil {
			return c.JSON(http.StatusOK, httpResponse.ParseError(err))
		}
		return c.JSON(http.StatusOK, httpResponse.NewRestResponse(http.StatusOK, constant.STATUS_MESSAGE_OK, 1))
	}
}

func (h Handler) Disable() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := utils.GetRequestCtx(c)
		user := c.Get("user").(*authModel.Response)
		if err := h.usecase.Disable(ctx, user.Id); err != nil {
			return c.JSON(http.StatusOK, httpResponse.ParseError(err))
		}
		return c.JSON(http.StatusOK, httpResponse.NewRestResponse(http.StatusOK, constant.STATUS_MESSAGE_OK, 1))
	}
}

// PromoteAdmin godoc
//
//	@Summary		Promote normal user to admin
//	@Description	Promote normal user to admin, guest not allowed
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			body	body	userModel.PromoteRequest	true	"body"
//	@Success		201
//	@Router			/users/promote/admin [post]
func (h Handler) PromoteAdmin() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := utils.GetRequestCtx(c)
		req := &userModel.PromoteRequest{}
		if err := utils.ReadBodyRequest(c, req); err != nil {
			log.Error().Err(err).Send()
			return c.JSON(http.StatusOK, httpResponse.NewInternalServerError(err))
		}
		// verify super admin
		user := c.Get("user").(*authModel.Response)
		if user.Email != h.cfg.Server.SuperAdmin {
			return c.JSON(http.StatusOK, httpResponse.NewRestResponse(http.StatusForbidden, constant.STATUS_CODE_FORBIDDEN, nil))
		}
		// promote admin
		if err := h.usecase.PromoteAdmin(ctx, user.Id, req.Email); err != nil {
			return c.JSON(http.StatusOK, httpResponse.ParseError(err))
		}
		return c.JSON(http.StatusOK, httpResponse.NewRestResponse(http.StatusOK, constant.STATUS_MESSAGE_OK, 1))
	}
}
