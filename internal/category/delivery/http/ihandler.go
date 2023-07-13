package handler

import "github.com/labstack/echo/v4"

// HTTP Handlers interface
type IHandler interface {
	MapRoutes(todoGroup *echo.Group)
	Create() echo.HandlerFunc
	GetListPaging() echo.HandlerFunc
}
