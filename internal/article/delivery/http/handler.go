package handler

import (
	"net/http"

	"github.com/Ho-Minh/InitiaRe-website/config"
	"github.com/Ho-Minh/InitiaRe-website/internal/article/models"
	"github.com/Ho-Minh/InitiaRe-website/internal/article/usecase"
	userModel "github.com/Ho-Minh/InitiaRe-website/internal/auth/models"
	"github.com/Ho-Minh/InitiaRe-website/internal/constants"
	"github.com/Ho-Minh/InitiaRe-website/internal/middleware"
	"github.com/Ho-Minh/InitiaRe-website/pkg/httpResponse"
	"github.com/Ho-Minh/InitiaRe-website/pkg/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type Handler struct {
	cfg     *config.Config
	usecase usecase.IUseCase
	mw      middleware.IMiddlewareManager
}

func NewHandler(cfg *config.Config, usecase usecase.IUseCase, mw middleware.IMiddlewareManager) IHandler {
	return Handler{
		cfg:     cfg,
		usecase: usecase,
		mw:      mw,
	}
}

// Map routes
func (h Handler) MapRoutes(group *echo.Group) {
	group.POST("", h.Create(), h.mw.AuthJWTMiddleware())
	group.GET("", h.GetListPaging())
}

// Create godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Create article
//	@Description	Create new article
//	@Tags			Article
//	@Accept			json
//	@Produce		json
//	@Param			body	body		models.CreateRequest	true	"body"
//	@Success		201		{object}	models.Response
//	@Router			/articles [post]
func (h Handler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := utils.GetRequestCtx(c)
		req := &models.CreateRequest{}
		if err := utils.ReadBodyRequest(c, req); err != nil {
			log.Error(err)
			return c.JSON(http.StatusOK, httpResponse.NewInternalServerError(err))
		}
		user := c.Get("user").(*userModel.Response)
		res, err := h.usecase.Create(ctx, user.Id, req.ToSaveRequest())
		if err != nil {
			return c.JSON(http.StatusOK, httpResponse.ParseError(err))
		}

		return c.JSON(http.StatusOK, httpResponse.NewRestResponse(http.StatusCreated, constants.STATUS_MESSAGE_CREATED, res))
	}
}

// GetListPaging godoc
//
//	@Summary		Get list article
//	@Description	Get list article with paging and filter
//	@Tags			Article
//	@Accept			json
//	@Produce		json
//	@Param			Page	query		int	true	"Page"
//	@Param			Size	query		int	true	"Size"
//	@Success		200		{object}	models.ListPaging
//	@Router			/articles [get]
func (h Handler) GetListPaging() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := utils.GetRequestCtx(c)
		req := &models.RequestList{}
		if err := utils.ReadQueryRequest(c, req); err != nil {
			log.Error(err)
			return c.JSON(http.StatusOK, httpResponse.NewInternalServerError(err))
		}

		res, err := h.usecase.GetListPaging(ctx, req)
		if err != nil {
			return c.JSON(http.StatusOK, httpResponse.ParseError(err))
		}

		return c.JSON(http.StatusOK, httpResponse.NewRestResponse(http.StatusOK, constants.STATUS_MESSAGE_OK, res))
	}
}
