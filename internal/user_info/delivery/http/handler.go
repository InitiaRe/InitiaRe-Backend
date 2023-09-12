package handler

import (
	"net/http"

	"github.com/Ho-Minh/InitiaRe-website/config"
	"github.com/Ho-Minh/InitiaRe-website/constant"
	userInfoModel "github.com/Ho-Minh/InitiaRe-website/internal/user_info/models"
	"github.com/Ho-Minh/InitiaRe-website/internal/user_info/usecase"
	"github.com/Ho-Minh/InitiaRe-website/pkg/httpResponse"
	"github.com/Ho-Minh/InitiaRe-website/pkg/utils"
	"github.com/rs/zerolog/log"

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
	group.POST("/enable", h.Enable())
	group.PUT("/disable", h.Disable())
}

func (h Handler) Enable() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := utils.GetRequestCtx(c)
		req := &userInfoModel.UserIdReq{}
		if err := utils.ReadBodyRequest(c, req); err != nil {
			log.Error().Err(err).Send()
			return c.JSON(http.StatusOK, httpResponse.NewInternalServerError(err))
		}
		res, err := h.usecase.Enable(ctx, req.Id)
		if err != nil {
			return c.JSON(http.StatusOK, httpResponse.ParseError(err))
		}
		return c.JSON(http.StatusOK, httpResponse.NewRestResponse(http.StatusOK, constant.STATUS_MESSAGE_OK, res))
	}
}

func (h Handler) Disable() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := utils.GetRequestCtx(c)
		req := &userInfoModel.UserIdReq{}
		if err := utils.ReadBodyRequest(c, req); err != nil {
			log.Error().Err(err).Send()
			return c.JSON(http.StatusOK, httpResponse.NewInternalServerError(err))
		}
		res, err := h.usecase.Disable(ctx, req.Id)
		if err != nil {
			return c.JSON(http.StatusOK, httpResponse.ParseError(err))
		}
		return c.JSON(http.StatusOK, httpResponse.NewRestResponse(http.StatusOK, constant.STATUS_MESSAGE_OK, res))
	}
}
