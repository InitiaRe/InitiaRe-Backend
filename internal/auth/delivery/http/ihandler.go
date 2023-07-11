package handler

import "github.com/labstack/echo/v4"

// Auth HTTP Handlers interface
type IHandler interface {
	MapAuthRoutes(authGroup *echo.Group)
	Login() echo.HandlerFunc
	Register() echo.HandlerFunc
}
