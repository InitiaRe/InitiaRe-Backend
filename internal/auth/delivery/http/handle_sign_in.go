package http

import (
	"net/http"

	"github.com/Ho-Minh/InitiaRe-website/pkg/httpResponse"
	"github.com/Ho-Minh/InitiaRe-website/pkg/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	// _ "github.com/Ho-Minh/InitiaRe-website/docs"
)

// SignIn godoc
//
//	@Summary		Sign in
//	@Description	Sign in and return token
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			Email		body		string	true	"Email"
//	@Param			Password	body		string	true	"Password"
//	@Success		200			{object}	models.User
//	@Router			/auth/sign-in [post]
func (h *authHandlers) SignIn() echo.HandlerFunc {
	type Login struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	return func(c echo.Context) error {
		ctx := utils.GetRequestCtx(c)
		login := &Login{}
		if err := utils.ReadRequest(c, login); err != nil {
			log.Error(err)
			return c.JSON(http.StatusOK, httpResponse.NewInternalServerError(err))
		}

		userWithToken, err := h.authUC.SignIn(ctx, login.Email, login.Password)
		if err != nil {
			return c.JSON(http.StatusOK, httpResponse.ParseError(err))
		}

		return c.JSON(http.StatusOK, httpResponse.NewRestResponse(http.StatusOK, "Success", userWithToken))
	}
}
