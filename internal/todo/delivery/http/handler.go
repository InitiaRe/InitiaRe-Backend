package handler

import (
	"net/http"
	"strings"

	"github.com/Ho-Minh/InitiaRe-website/config"
	userModel "github.com/Ho-Minh/InitiaRe-website/internal/auth/models"
	"github.com/Ho-Minh/InitiaRe-website/internal/constants"
	"github.com/Ho-Minh/InitiaRe-website/internal/middleware"
	"github.com/Ho-Minh/InitiaRe-website/internal/todo/models"
	"github.com/Ho-Minh/InitiaRe-website/internal/todo/usecase"
	"github.com/Ho-Minh/InitiaRe-website/pkg/httpResponse"
	"github.com/Ho-Minh/InitiaRe-website/pkg/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type Handler struct {
	cfg     *config.Config
	usecase usecase.IUseCase
	mw      *middleware.MiddlewareManager
}

func NewHandler(cfg *config.Config, usecase usecase.IUseCase, mw *middleware.MiddlewareManager) IHandler {
	return Handler{
		cfg:     cfg,
		usecase: usecase,
		mw:      mw,
	}
}

// Map todo routes
func (h Handler) MapTodoRoutes(todoGroup *echo.Group) {
	todoGroup.POST("", h.Create(), h.mw.AuthJWTMiddleware())
	todoGroup.GET("", h.GetListPaging(), h.mw.AuthJWTMiddleware())
}

// Create godoc
//
//	@Summary		Create todo
//	@Description	Create new todo
//	@Tags			Todo
//	@Accept			json
//	@Produce		json
//	@Param			Content	body		string	true	"Content"
//	@Success		201		{object}	models.Response
//	@Router			/todo [post]
func (h Handler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := utils.GetRequestCtx(c)
		todo := &models.SaveRequest{}
		if err := utils.ReadBodyRequest(c, todo); err != nil {
			log.Error(err)
			return c.JSON(http.StatusOK, httpResponse.NewInternalServerError(err))
		}
		user := c.Get("user").(*userModel.Response)
		createdTodo, err := h.usecase.Create(ctx, user.Id, todo)
		if err != nil {
			if strings.Contains(err.Error(), constants.STATUS_CODE_BAD_REQUEST) {
				return c.JSON(http.StatusOK, httpResponse.NewBadRequestError(utils.GetErrorMessage(err)))
			} else {
				return c.JSON(http.StatusOK, httpResponse.NewInternalServerError(err))
			}
		}

		return c.JSON(http.StatusOK, httpResponse.NewRestResponse(http.StatusCreated, constants.STATUS_MESSAGE_CREATED, createdTodo))
	}
}

// GetListPaging godoc
//
//	@Summary		Get list todo
//	@Description	Get list todo with paging and filter
//	@Tags			Todo
//	@Accept			json
//	@Produce		json
//	@Param			Page	query		int	true	"Page"
//	@Param			Size	query		int	true	"Size"
//	@Success		200		{object}	models.ListPaging
//	@Router			/todo [get]
func (h Handler) GetListPaging() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := utils.GetRequestCtx(c)
		req := &models.RequestList{}
		if err := utils.ReadQueryRequest(c, req); err != nil {
			log.Error(err)
			return c.JSON(http.StatusOK, httpResponse.NewInternalServerError(err))
		}

		listTodo, err := h.usecase.GetListPaging(ctx, req)
		if err != nil {
			if strings.Contains(err.Error(), constants.STATUS_CODE_BAD_REQUEST) {
				return c.JSON(http.StatusOK, httpResponse.NewBadRequestError(utils.GetErrorMessage(err)))
			} else {
				return c.JSON(http.StatusOK, httpResponse.NewInternalServerError(err))
			}
		}

		return c.JSON(http.StatusOK, httpResponse.NewRestResponse(http.StatusOK, constants.STATUS_MESSAGE_OK, listTodo))
	}
}
