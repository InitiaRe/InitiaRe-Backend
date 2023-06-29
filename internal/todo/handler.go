package todo

import "github.com/labstack/echo/v4"

// Auth HTTP Handlers interface
type Handlers interface {
	Create() echo.HandlerFunc
	// Update() echo.HandlerFunc
	// GetByID() echo.HandlerFunc
	// Delete() echo.HandlerFunc
	// GetTodos() echo.HandlerFunc
	// SearchByTitle() echo.HandlerFunc
}
