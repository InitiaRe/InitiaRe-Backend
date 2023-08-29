package handler

import "github.com/labstack/echo/v4"

// HTTP Handlers interface
type IHandler interface {
	MapRoutes(group *echo.Group)
	Disable() echo.HandlerFunc
	Enable() echo.HandlerFunc
}
