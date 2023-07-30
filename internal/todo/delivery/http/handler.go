package handler

import (
	"net/http"
	"strconv"
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
func (h Handler) MapRoutes(todoGroup *echo.Group) {
	todoGroup.POST("", h.Create(), h.mw.AuthJWTMiddleware())
	todoGroup.PUT("/:id", h.Update(), h.mw.AuthJWTMiddleware())
	todoGroup.DELETE("/:id", h.Delete(), h.mw.AuthJWTMiddleware())
	todoGroup.GET("", h.GetListPaging(), h.mw.AuthJWTMiddleware())
}

// Create godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Create todo
//	@Description	Create new todo
//	@Tags			Todo
//	@Accept			json
//	@Produce		json
//	@Param			body	body		models.CreateRequest	true	"body"
//	@Success		201		{object}	models.Response
//	@Router			/todos [post]
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
			if strings.Contains(err.Error(), constants.STATUS_CODE_BAD_REQUEST) {
				return c.JSON(http.StatusOK, httpResponse.NewBadRequestError(utils.GetErrorMessage(err)))
			} else {
				return c.JSON(http.StatusOK, httpResponse.NewInternalServerError(err))
			}
		}

		return c.JSON(http.StatusOK, httpResponse.NewRestResponse(http.StatusCreated, constants.STATUS_MESSAGE_CREATED, res))
	}
}

// Update godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Update todo
//	@Description	Update todo
//	@Tags			Todo
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int						true	"Id"
//	@Param			body	body		models.UpdateRequest	true	"body"
//	@Success		200		{object}	models.Response
//	@Router			/todos/{id} [put]
func (h Handler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := utils.GetRequestCtx(c)
		req := &models.UpdateRequest{}
		if err := utils.ReadBodyRequest(c, req); err != nil {
			log.Error(err)
			return c.JSON(http.StatusOK, httpResponse.NewInternalServerError(err))
		}
		user := c.Get("user").(*userModel.Response)
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Error(err)
			return c.JSON(http.StatusOK, httpResponse.NewInternalServerError(err))
		}
		res, err := h.usecase.Update(ctx, user.Id, req.ToSaveRequest(id))
		if err != nil {
			if strings.Contains(err.Error(), constants.STATUS_CODE_BAD_REQUEST) {
				return c.JSON(http.StatusOK, httpResponse.NewBadRequestError(utils.GetErrorMessage(err)))
			} else {
				return c.JSON(http.StatusOK, httpResponse.NewInternalServerError(err))
			}
		}

		return c.JSON(http.StatusOK, httpResponse.NewRestResponse(http.StatusOK, constants.STATUS_MESSAGE_OK, res))
	}
}

// Delete godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Delete todo
//	@Description	Delete todo
//	@Tags			Todo
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Id"
//	@Success		200	{object}	models.Response
//	@Router			/todos/{id} [delete]
func (h Handler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := utils.GetRequestCtx(c)
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Error(err)
			return c.JSON(http.StatusOK, httpResponse.NewInternalServerError(err))
		}
		res, err := h.usecase.Delete(ctx,id)
		if err != nil {
			if strings.Contains(err.Error(), constants.STATUS_CODE_BAD_REQUEST) {
				return c.JSON(http.StatusOK, httpResponse.NewBadRequestError(utils.GetErrorMessage(err)))
			} else {
				return c.JSON(http.StatusOK, httpResponse.NewInternalServerError(err))
			}
		}

		return c.JSON(http.StatusOK, httpResponse.NewRestResponse(http.StatusOK, constants.STATUS_MESSAGE_OK, res))
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
//	@Router			/todos [get]
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
			if strings.Contains(err.Error(), constants.STATUS_CODE_BAD_REQUEST) {
				return c.JSON(http.StatusOK, httpResponse.NewBadRequestError(utils.GetErrorMessage(err)))
			} else {
				return c.JSON(http.StatusOK, httpResponse.NewInternalServerError(err))
			}
		}

		return c.JSON(http.StatusOK, httpResponse.NewRestResponse(http.StatusOK, constants.STATUS_MESSAGE_OK, res))
	}
}
