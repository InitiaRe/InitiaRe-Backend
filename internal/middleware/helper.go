package middleware

import (
	"errors"

	"github.com/Ho-Minh/InitiaRe-website/internal/constants"

	"github.com/Ho-Minh/InitiaRe-website/pkg/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func (mw *middlewareManager) validateJWTToken(c echo.Context, tokenString string) error {
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

		user, err := mw.authRedisRepo.GetById(c.Request().Context(), utils.GenerateUserKey(int(userId)))
		if err != nil {
			return err
		}
		c.Set("user", user.Export())
	}
	return nil
}
