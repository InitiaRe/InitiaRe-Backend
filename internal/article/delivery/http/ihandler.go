package handler

import "github.com/labstack/echo/v4"

// HTTP Handlers interface
type IHandler interface {
	MapRoutes(group *echo.Group)
	Create() echo.HandlerFunc
	GetListPaging() echo.HandlerFunc
	GetApprovedArticle() echo.HandlerFunc
	GetById() echo.HandlerFunc
	Update() echo.HandlerFunc
	GetByMe() echo.HandlerFunc
	Approve() echo.HandlerFunc
	Disable() echo.HandlerFunc
}
