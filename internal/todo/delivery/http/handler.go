package http

import (
	"github.com/Ho-Minh/InitiaRe-website/config"
	"github.com/Ho-Minh/InitiaRe-website/internal/todo"
)

type todoHandlers struct {
	cfg    *config.Config
	todoUC todo.UseCase
}

func NewTodoHandlers(cfg *config.Config, todoUC todo.UseCase) todo.Handlers {
	return &todoHandlers{
		cfg:    cfg,
		todoUC: todoUC,
	}
}
