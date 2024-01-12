package handler

import "github.com/labstack/echo/v4"

// HTTP Handlers interface
type IHandler interface {
	MapRoutes(group *echo.Group)
	GetRating() echo.HandlerFunc
	Vote() echo.HandlerFunc
}
