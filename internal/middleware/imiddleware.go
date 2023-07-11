package middleware

import "github.com/labstack/echo/v4"

type IMiddlewareManager interface {
	AuthJWTMiddleware() echo.MiddlewareFunc
}
