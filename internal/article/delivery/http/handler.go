package handler

import (
	"net/http"
	"strconv"
	"strings"

	"InitiaRe-website/config"
	"InitiaRe-website/constant"
	"InitiaRe-website/internal/article/models"
	"InitiaRe-website/internal/article/usecase"
	userModel "InitiaRe-website/internal/auth/models"
	"InitiaRe-website/internal/middleware"
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
	group.POST("", h.Create(), h.mw.AuthJWTMiddleware())

	group.GET("/:id", h.GetById())
	group.GET("/approved-article", h.GetApprovedArticle())
	group.GET("", h.GetListPaging())

	group.PUT("/:id", h.Update(), h.mw.AuthJWTMiddleware())
	group.GET("/me", h.GetByMe(), h.mw.AuthJWTMiddleware())

	// Admin role
	group.POST("/approve", h.Approve(), h.mw.AuthJWTMiddleware())
	group.POST("/disable", h.Disable(), h.mw.AuthJWTMiddleware())
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
			log.Error().Err(err).Send()
			return c.JSON(http.StatusOK, httpResponse.NewInternalServerError(err))
		}
		if req.TypeId <= 0 {
			return c.JSON(http.StatusBadRequest, httpResponse.NewBadRequestError("Invalid TypeId"))
		}
		user := c.Get("user").(*userModel.Response)
		res, err := h.usecase.Create(ctx, user.Id, req.ToSaveRequest())
		if err != nil {
			return c.JSON(http.StatusOK, httpResponse.ParseError(err))
		}

		return c.JSON(http.StatusOK, httpResponse.NewRestResponse(http.StatusCreated, constant.STATUS_MESSAGE_CREATED, res))
	}
}

// GetListPaging godoc
//
//	@Summary		Get list article
//	@Description	Get list article with paging and filter
//	@Tags			Article
//	@Accept			json
//	@Produce		json
//	@Param			title			query		string	false	"Title"
//	@Param			status_id		query		int		false	"Status"
//	@Param			type_id			query		int		false	"Type"
//	@Param			category_id		query		int		false	"Category"
//	@Param			category_ids	query		string	false	"Category"
//	@Param			email			query		string	false	"Email"
//	@Param			page			query		int		true	"Page"
//	@Param			size			query		int		true	"Size"
//	@Success		200				{object}	models.ListPaging
//	@Router			/articles [get]
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

// GetApprovedArticle godoc
//
//	@Summary		Get approved article
//	@Description	Get the list of approved articles
//	@Tags			Article
//	@Accept			json
//	@Produce		json
//	@Param			title			query		string	false	"Title"
//	@Param			type_id			query		int		false	"Type"
//	@Param			category_id		query		int		false	"Category"
//	@Param			category_ids	query		string	false	"Category"
//	@Param			email			query		string	false	"Email"
//	@Param			page			query		int		true	"Page"
//	@Param			size			query		int		true	"Size"
//	@Success		200				{object}	models.ListPaging
//	@Router			/articles/approved-article [get]
func (h Handler) GetApprovedArticle() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := utils.GetRequestCtx(c)
		req := &models.RequestList{}
		if err := utils.ReadQueryRequest(c, req); err != nil {
			log.Error().Err(err).Send()
			return c.JSON(http.StatusOK, httpResponse.NewInternalServerError(err))
		}

		res, err := h.usecase.GetApprovedArticle(ctx, req)
		if err != nil {
			return c.JSON(http.StatusOK, httpResponse.ParseError(err))
		}

		return c.JSON(http.StatusOK, httpResponse.NewRestResponse(http.StatusOK, constant.STATUS_MESSAGE_OK, res))
	}
}

// Update godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Update article
//	@Description	Update an existing article
//	@Tags			Article
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int						true	"Id"
//	@Param			body	body		models.UpdateRequest	true	"body"
//	@Success		200		{object}	models.Response
//	@Router			/articles/{id} [put]
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
			if strings.Contains(err.Error(), constant.STATUS_CODE_BAD_REQUEST) {
				return c.JSON(http.StatusOK, httpResponse.NewBadRequestError(utils.GetErrorMessage(err)))
			} else {
				return c.JSON(http.StatusOK, httpResponse.NewInternalServerError(err))
			}
		}

		return c.JSON(http.StatusOK, httpResponse.NewRestResponse(http.StatusOK, constant.STATUS_MESSAGE_OK, res))
	}
}

// GetById godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Get detail article
//	@Description	Get detail article
//	@Tags			Article
//	@Produce		json
//	@Param			id	path		int	true	"Id"
//	@Success		200	{object}	models.Response
//	@Router			/articles/{id} [get]
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

// GetByMe godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Get list current user articles
//	@Description	Get list current user articles by token
//	@Tags			Article
//	@Produce		json
//	@Param			title			query		string	false	"Title"
//	@Param			status_id		query		int		false	"Status"
//	@Param			type_id			query		int		false	"Type"
//	@Param			category_id		query		int		false	"Category"
//	@Param			category_ids	query		string	false	"Category"
//	@Param			page			query		int		true	"Page"
//	@Param			size			query		int		true	"Size"
//	@Success		200				{object}	models.Response
//	@Router			/articles/me [get]
func (h Handler) GetByMe() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := utils.GetRequestCtx(c)
		req := &models.RequestList{}
		if err := utils.ReadQueryRequest(c, req); err != nil {
			log.Error().Err(err).Send()
			return c.JSON(http.StatusOK, httpResponse.NewInternalServerError(err))
		}

		user := c.Get("user")
		req.CreatedBy = user.(*userModel.Response).Id
		res, err := h.usecase.GetListPaging(ctx, req)
		if err != nil {
			return c.JSON(http.StatusOK, httpResponse.ParseError(err))
		}

		return c.JSON(http.StatusOK, httpResponse.NewRestResponse(http.StatusOK, constant.STATUS_MESSAGE_OK, res))
	}
}

// Approve godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Approve article
//	@Description	Approve article by Id
//	@Tags			Article
//	@Accept			json
//	@Produce		json
//	@Param			body	body		models.ApproveRequest	true	"body"
//	@Success		200		{object}	httpResponse.RestResponse
//	@Router			/articles/approve [post]
func (h Handler) Approve() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := utils.GetRequestCtx(c)
		req := &models.ApproveRequest{}
		if err := utils.ReadBodyRequest(c, req); err != nil {
			log.Error().Err(err).Send()
			return c.JSON(http.StatusOK, httpResponse.NewInternalServerError(err))
		}
		if err := h.usecase.Approve(ctx, req.Id); err != nil {
			return c.JSON(http.StatusOK, httpResponse.ParseError(err))
		}
		return c.JSON(http.StatusOK, httpResponse.NewRestResponse(http.StatusOK, constant.STATUS_MESSAGE_OK, 1))
	}
}

// Disable godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Disable article
//	@Description	Disable article by Id
//	@Tags			Article
//	@Accept			json
//	@Produce		json
//	@Param			body	body		models.DisableRequest	true	"body"
//	@Success		200		{object}	httpResponse.RestResponse
//	@Router			/articles/disable [post]
func (h Handler) Disable() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := utils.GetRequestCtx(c)
		req := &models.DisableRequest{}
		if err := utils.ReadBodyRequest(c, req); err != nil {
			log.Error().Err(err).Send()
			return c.JSON(http.StatusOK, httpResponse.NewInternalServerError(err))
		}
		if err := h.usecase.Disable(ctx, req.Id); err != nil {
			return c.JSON(http.StatusOK, httpResponse.ParseError(err))
		}
		return c.JSON(http.StatusOK, httpResponse.NewRestResponse(http.StatusOK, constant.STATUS_MESSAGE_OK, 1))
	}
}
