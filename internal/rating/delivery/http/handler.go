package handler

import (
	"net/http"
	"strconv"

	"github.com/Ho-Minh/InitiaRe-website/config"
	"github.com/Ho-Minh/InitiaRe-website/constant"
	userModel "github.com/Ho-Minh/InitiaRe-website/internal/auth/models"
	"github.com/Ho-Minh/InitiaRe-website/internal/middleware"
	"github.com/Ho-Minh/InitiaRe-website/internal/rating/models"
	"github.com/Ho-Minh/InitiaRe-website/internal/rating/usecase"
	"github.com/Ho-Minh/InitiaRe-website/pkg/httpResponse"
	"github.com/Ho-Minh/InitiaRe-website/pkg/utils"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
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
}

func (h Handler) Vote() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := utils.GetRequestCtx(c)
		req := &models.VoteRequest{}
		if err := utils.ReadBodyRequest(c, req); err != nil {
			log.Error().Err(err).Send()
			return c.JSON(http.StatusOK, httpResponse.NewInternalServerError(err))
		}

		user := c.Get("user").(*userModel.Response)
		res, err := h.usecase.Vote(ctx, req.ToSaveRequest(user.Id))
		if err != nil {
			return c.JSON(http.StatusOK, httpResponse.ParseError(err))
		}

		return c.JSON(http.StatusOK, httpResponse.NewRestResponse(http.StatusOK, constant.STATUS_MESSAGE_OK, res))
	}
}

func (h Handler) GetRating() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := utils.GetRequestCtx(c)
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Error().Err(err).Send()
			return c.JSON(http.StatusOK, httpResponse.NewInternalServerError(err))
		}

		res, err := h.usecase.GetRating(ctx, id)
		if err != nil {
			return c.JSON(http.StatusOK, httpResponse.ParseError(err))
		}

		return c.JSON(http.StatusOK, httpResponse.NewRestResponse(http.StatusOK, constant.STATUS_MESSAGE_OK, res))
	}
}
