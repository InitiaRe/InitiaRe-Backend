package handler

import "github.com/labstack/echo/v4"

// HTTP Handlers interface
type IHandler interface {
	MapRoutes(group *echo.Group)
	Login() echo.HandlerFunc
	Register() echo.HandlerFunc
	ResetPassword() echo.HandlerFunc
}
