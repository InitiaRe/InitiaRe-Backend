package middleware

import (
	"net/http"
	"strings"

	"github.com/Ho-Minh/InitiaRe-website/config"
	authRepo "github.com/Ho-Minh/InitiaRe-website/internal/auth/repository"
	"github.com/Ho-Minh/InitiaRe-website/pkg/httpResponse"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type middlewareManager struct {
	cfg      *config.Config
	authRedisRepo authRepo.IRedisRepository
}

func NewMiddlewareManager(cfg *config.Config, authRedisRepo authRepo.IRedisRepository) IMiddlewareManager {
	return &middlewareManager{
		cfg:      cfg,
		authRedisRepo: authRedisRepo,
	}
}

// JWT way of auth using Authorization header
func (mw *middlewareManager) AuthJWTMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			bearerHeader := c.Request().Header.Get("Authorization")

			if bearerHeader != "" {
				log.Infof("auth middleware bearerHeader %s", bearerHeader)
				headerParts := strings.Split(bearerHeader, " ")
				if len(headerParts) != 2 {
					log.Errorf("auth middleware: length header invalid")
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
