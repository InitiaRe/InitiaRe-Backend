package middleware

import (
	"errors"

	"InitiaRe-website/constant"
	"InitiaRe-website/pkg/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

func (mw *middlewareManager) validateJWTToken(c echo.Context, tokenString string) error {
	if tokenString == "" {
		return errors.New(constant.STATUS_MESSAGE_INVALID_JWT_TOKEN)
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Error().Msgf("unexpected signin method %v", token.Header["alg"])
			return nil, utils.NewError(constant.STATUS_CODE_INTERNAL_SERVER, constant.STATUS_MESSAGE_INTERNAL_SERVER_ERROR)
		}
		return []byte(mw.cfg.Auth.Secret), nil
	})
	if err != nil {
		log.Error().Err(err).Str("service", "jwt.Parse").Send()
		return err
	}
	if !token.Valid {
		return errors.New(constant.STATUS_MESSAGE_INVALID_JWT_TOKEN)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		userId, ok := claims["Id"].(float64)
		if !ok {
			return errors.New(constant.STATUS_MESSAGE_INVALID_JWT_TOKEN)
		}

		user, err := mw.cacheRepo.GetById(c.Request().Context(), utils.GenerateUserKey(int(userId)))
		if err != nil {
			return err
		}
		c.Set("user", user.Export())
	}
	return nil
}
