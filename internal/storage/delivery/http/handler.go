package handler

import (
	"net/http"

	"InitiaRe-website/config"
	"InitiaRe-website/constant"
	userModel "InitiaRe-website/internal/auth/models"
	"InitiaRe-website/internal/middleware"
	"InitiaRe-website/internal/storage/models"
	"InitiaRe-website/internal/storage/usecase"
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
	group.POST("/media/upload", h.UploadMedia(), h.mw.AuthJWTMiddleware())
}

// Create godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Upload media
//	@Description	Upload media file (pdf, docs, images, videos, etc.)
//	@Tags			Storage
//	@Accept			mpfd
//	@Produce		json
//	@Param			file	formData	file	true	"binary file"
//	@Success		201		{object}	models.Response
//	@Router			/storage/media/upload [post]
func (h Handler) UploadMedia() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := utils.GetRequestCtx(c)
		file, err := c.FormFile("file")
		if err != nil {
			log.Error().Err(err).Send()
			return c.JSON(http.StatusOK, httpResponse.ParseError(err))
		}
		req := models.UploadRequest{
			File: file,
		}
		user := c.Get("user").(*userModel.Response)
		res, err := h.usecase.UploadMedia(ctx, user.Id, &req)
		if err != nil {
			return c.JSON(http.StatusOK, httpResponse.ParseError(err))
		}

		return c.JSON(http.StatusOK, httpResponse.NewRestResponse(http.StatusCreated, constant.STATUS_MESSAGE_CREATED, res))
	}
}
