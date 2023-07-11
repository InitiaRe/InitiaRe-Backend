package handler

import (
	"net/http"
	"strings"

	"github.com/Ho-Minh/InitiaRe-website/config"
	"github.com/Ho-Minh/InitiaRe-website/internal/auth/models"
	"github.com/Ho-Minh/InitiaRe-website/internal/auth/usecase"
	"github.com/Ho-Minh/InitiaRe-website/internal/constants"
	"github.com/Ho-Minh/InitiaRe-website/pkg/httpResponse"
	"github.com/Ho-Minh/InitiaRe-website/pkg/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type Handler struct {
	cfg     *config.Config
	usecase usecase.IUseCase
}

func NewHandler(cfg *config.Config, usecase usecase.IUseCase) IHandler {
	return Handler{
		cfg:     cfg,
		usecase: usecase,
	}
}

// Map routes
func (h Handler) MapRoutes(group *echo.Group) {
	group.POST("/register", h.Register())
	group.POST("/login", h.Login())
}

// Login godoc
//
//	@Summary		Login
//	@Description	Login and return token
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			Email		body		string	true	"Email"
//	@Param			Password	body		string	true	"Password"
//	@Success		200			{object}	models.UserWithToken
//	@Router			/auth/login [post]
func (h Handler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := utils.GetRequestCtx(c)
		req := &models.LoginRequest{}
		if err := utils.ReadBodyRequest(c, req); err != nil {
			log.Error(err)
			return c.JSON(http.StatusOK, httpResponse.NewInternalServerError(err))
		}

		res, err := h.usecase.Login(ctx, req)
		if err != nil {
			return c.JSON(http.StatusOK, httpResponse.ParseError(err))
		}

		return c.JSON(http.StatusOK, httpResponse.NewRestResponse(http.StatusOK, "Success", res))
	}
}

// Register godoc
//
//	@Summary		Create new user
//	@Description	Create new user, returns user and token
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			FirstName	body		string	true	"First name"
//	@Param			LastName	body		string	true	"Last name"
//	@Param			Email		body		string	true	"Email"
//	@Param			Password	body		string	true	"Password"
//	@Param			Gender		body		string	true	"Gender"
//	@Param			School		body		string	false	"School"
//	@Param			Birthday	body		string	false	"Gender"
//	@Success		201			{object}	models.Response
//	@Router			/auth/register [post]
func (h Handler) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := utils.GetRequestCtx(c)
		req := &models.SaveRequest{}
		if err := utils.ReadBodyRequest(c, req); err != nil {
			log.Error(err)
			return c.JSON(http.StatusOK, httpResponse.NewInternalServerError(err))
		}

		res, err := h.usecase.Register(ctx, req)
		if err != nil {
			if strings.Contains(err.Error(), constants.STATUS_CODE_BAD_REQUEST) {
				return c.JSON(http.StatusOK, httpResponse.NewBadRequestError(utils.GetErrorMessage(err)))
			} else {
				return c.JSON(http.StatusOK, httpResponse.NewInternalServerError(err))
			}
		}
		return c.JSON(http.StatusCreated, httpResponse.NewRestResponse(http.StatusCreated, constants.STATUS_MESSAGE_CREATED, res))
	}
}
