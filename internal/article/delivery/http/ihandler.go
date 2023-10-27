package handler

import "github.com/labstack/echo/v4"

// HTTP Handlers interface
type IHandler interface {
	MapRoutes(group *echo.Group)
	Create() echo.HandlerFunc
	GetListPaging() echo.HandlerFunc
	GetById() echo.HandlerFunc
	Update() echo.HandlerFunc
	GetByMe() echo.HandlerFunc
	ApproveArticle() echo.HandlerFunc
}
