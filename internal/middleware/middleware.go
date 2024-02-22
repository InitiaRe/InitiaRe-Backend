package middleware

import (
	"net/http"
	"strings"

	"InitiaRe-website/config"
	authRepo "InitiaRe-website/internal/auth/repository"
	"InitiaRe-website/pkg/httpResponse"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

type middlewareManager struct {
	cfg       *config.Config
	cacheRepo authRepo.ICacheRepository
}

func NewMiddlewareManager(cfg *config.Config, cacheRepo authRepo.ICacheRepository) IMiddlewareManager {
	return &middlewareManager{
		cfg:       cfg,
		cacheRepo: cacheRepo,
	}
}

// JWT way of auth using Authorization header
func (mw *middlewareManager) AuthJWTMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			bearerHeader := c.Request().Header.Get("Authorization")

			if bearerHeader != "" {
				// log.Info().Msgf("Middleware bearerHeader: %s", bearerHeader)
				headerParts := strings.Split(bearerHeader, " ")
				if len(headerParts) != 2 {
					log.Error().Msg("auth middleware: length header invalid")
					return c.JSON(http.StatusOK, httpResponse.NewUnauthorizedError(nil))
				}
				tokenString := headerParts[1]

				if err := mw.validateJWTToken(c, tokenString); err != nil {
					log.Error().Err(err).Str("service", "mw.validateJWTToken").Send()
					return c.JSON(http.StatusUnauthorized, httpResponse.NewUnauthorizedError(nil))
				}

				return next(c)
			} else {
				log.Error().Msg("Invalid Authorization header")
				return c.JSON(http.StatusOK, httpResponse.NewUnauthorizedError(nil))
			}
		}
	}
}
