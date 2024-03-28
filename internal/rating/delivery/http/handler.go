package handler

import (
	"net/http"
	"strconv"

	"InitiaRe-website/config"
	"InitiaRe-website/constant"
	userModel "InitiaRe-website/internal/auth/models"
	"InitiaRe-website/internal/middleware"
	"InitiaRe-website/internal/rating/models"
	"InitiaRe-website/internal/rating/usecase"
	"InitiaRe-website/pkg/httpResponse"
	"InitiaRe-website/pkg/utils"

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
	group.POST("/vote/:id", h.Vote(), h.mw.AuthJWTMiddleware())
	group.GET("/vote/:id", h.GetRating())
}

func (h Handler) Vote() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := utils.GetRequestCtx(c)
		req := &models.VoteRequest{}
		if err := utils.ReadBodyRequest(c, req); err != nil {
			log.Error().Err(err).Send()
			return c.JSON(http.StatusOK, httpResponse.NewInternalServerError(err))
		}

		articleId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Error().Err(err).Send()
			return c.JSON(http.StatusOK, httpResponse.NewInternalServerError(err))
		}
		req.ArticleId = articleId

		user := c.Get("user").(*userModel.Response)
		err = h.usecase.Vote(ctx, req.ToSaveRequest(user.Id))
		if err != nil {
			return c.JSON(http.StatusOK, httpResponse.ParseError(err))
		}

		return c.JSON(http.StatusOK, httpResponse.NewRestResponse(http.StatusOK, constant.STATUS_MESSAGE_OK, nil))
	}
}

// GetRating godoc
//
//	@Summary		Get article vote
//	@Description	Get article vote
//	@Tags			Rating
//	@Produce		json
//	@Success		200
//	@Router			/rating/:id [get]
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
