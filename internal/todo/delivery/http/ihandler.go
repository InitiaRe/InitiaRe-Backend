package handler

import "github.com/labstack/echo/v4"

// Auth HTTP Handlers interface
type IHandler interface {
	MapTodoRoutes(todoGroup *echo.Group)
	Create() echo.HandlerFunc
	GetListPaging() echo.HandlerFunc
}
