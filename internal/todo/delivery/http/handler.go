package handler

import (
	"net/http"
	"strconv"

	"github.com/Ho-Minh/InitiaRe-website/config"
	"github.com/Ho-Minh/InitiaRe-website/constant"
	userModel "github.com/Ho-Minh/InitiaRe-website/internal/auth/models"
	"github.com/Ho-Minh/InitiaRe-website/internal/middleware"
	"github.com/Ho-Minh/InitiaRe-website/internal/todo/models"
	"github.com/Ho-Minh/InitiaRe-website/internal/todo/usecase"
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
	group.POST("", h.Create(), h.mw.AuthJWTMiddleware())
	group.PUT("/:id", h.Update(), h.mw.AuthJWTMiddleware())
	group.DELETE("/:id", h.Delete(), h.mw.AuthJWTMiddleware())
	group.GET("", h.GetListPaging(), h.mw.AuthJWTMiddleware())
	group.GET("/:id", h.GetById(), h.mw.AuthJWTMiddleware())
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
			log.Error().Err(err).Send()
			return c.JSON(http.StatusOK, httpResponse.NewInternalServerError(err))
		}
		user := c.Get("user").(*userModel.Response)
		res, err := h.usecase.Create(ctx, user.Id, req.ToSaveRequest())
		if err != nil {
			return c.JSON(http.StatusOK, httpResponse.ParseError(err))
		}

		return c.JSON(http.StatusOK, httpResponse.NewRestResponse(http.StatusCreated, constant.STATUS_MESSAGE_CREATED, res))
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
			log.Error().Err(err).Send()
			return c.JSON(http.StatusOK, httpResponse.NewInternalServerError(err))
		}
		user := c.Get("user").(*userModel.Response)
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Error().Err(err).Send()
			return c.JSON(http.StatusOK, httpResponse.NewInternalServerError(err))
		}
		res, err := h.usecase.Update(ctx, user.Id, req.ToSaveRequest(id))
		if err != nil {
			return c.JSON(http.StatusOK, httpResponse.ParseError(err))
		}

		return c.JSON(http.StatusOK, httpResponse.NewRestResponse(http.StatusOK, constant.STATUS_MESSAGE_OK, res))
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
			log.Error().Err(err).Send()
			return c.JSON(http.StatusOK, httpResponse.NewInternalServerError(err))
		}
		user := c.Get("user").(*userModel.Response)
		res, err := h.usecase.Delete(ctx, user.Id, id)
		if err != nil {
			return c.JSON(http.StatusOK, httpResponse.ParseError(err))
		}

		return c.JSON(http.StatusOK, httpResponse.NewRestResponse(http.StatusOK, constant.STATUS_MESSAGE_OK, res))
	}
}

// GetListPaging godoc
//
//	@Security		ApiKeyAuth
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
			log.Error().Err(err).Send()
			return c.JSON(http.StatusOK, httpResponse.NewInternalServerError(err))
		}

		res, err := h.usecase.GetListPaging(ctx, req)
		if err != nil {
			return c.JSON(http.StatusOK, httpResponse.ParseError(err))
		}

		return c.JSON(http.StatusOK, httpResponse.NewRestResponse(http.StatusOK, constant.STATUS_MESSAGE_OK, res))
	}
}

// GetById godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Get detail todo
//	@Description	Get detail todo
//	@Tags			Todo
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Id"
//	@Success		200	{object}	models.Response
//	@Router			/todos/{id} [get]
func (h Handler) GetById() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := utils.GetRequestCtx(c)
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Error().Err(err).Send()
			return c.JSON(http.StatusOK, httpResponse.NewInternalServerError(err))
		}

		res, err := h.usecase.GetById(ctx, id)
		if err != nil {
			return c.JSON(http.StatusOK, httpResponse.ParseError(err))
		}

		return c.JSON(http.StatusOK, httpResponse.NewRestResponse(http.StatusOK, constant.STATUS_MESSAGE_OK, res))
	}
}
