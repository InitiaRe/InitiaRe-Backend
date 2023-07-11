package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/Ho-Minh/InitiaRe-website/internal/constants"
	"github.com/Ho-Minh/InitiaRe-website/pkg/httpResponse"

	"github.com/Ho-Minh/InitiaRe-website/pkg/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

// JWT way of auth using Authorization header
func (mw *MiddlewareManager) AuthJWTMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			bearerHeader := c.Request().Header.Get("Authorization")

			if bearerHeader != "" {
				log.Infof("auth middleware bearerHeader %s", bearerHeader)
				headerParts := strings.Split(bearerHeader, " ")
				if len(headerParts) != 2 {
					log.Errorf("auth middleware: %s", len(headerParts) != 2)
					return c.JSON(http.StatusOK, httpResponse.NewUnauthorizedError(nil))
				}
				tokenString := headerParts[1]

				if err := mw.validateJWTToken(c, tokenString); err != nil {
					log.Errorf("middleware validateJWTToken: %s", err.Error())
					return c.JSON(http.StatusUnauthorized, httpResponse.NewUnauthorizedError(nil))
				}

				return next(c)
			} else {
				log.Errorf("Invalid Authorization header")
				return c.JSON(http.StatusOK, httpResponse.NewUnauthorizedError(nil))
			}
		}
	}
}

func (mw *MiddlewareManager) validateJWTToken(c echo.Context, tokenString string) error {
	if tokenString == "" {
		return errors.New(constants.STATUS_MESSAGE_INVALID_JWT_TOKEN)
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Errorf("unexpected signin method %v", token.Header["alg"])
			return nil, utils.NewError(constants.STATUS_CODE_INTERNAL_SERVER, constants.STATUS_MESSAGE_INTERNAL_SERVER_ERROR)
		}
		return []byte(mw.cfg.Auth.JWTSecret), nil
	})
	if err != nil {
		log.Errorf("jwt.Parse %v", err)
		return err
	}
	if !token.Valid {
		return errors.New(constants.STATUS_MESSAGE_INVALID_JWT_TOKEN)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		userId, ok := claims["Id"].(float64)
		if !ok {
			return errors.New(constants.STATUS_MESSAGE_INVALID_JWT_TOKEN)
		}

		user, err := mw.authRepo.GetById(c.Request().Context(), int(userId))
		if err != nil {
			return err
		}

		c.Set("user", user.Export())
	}
	return nil
}
