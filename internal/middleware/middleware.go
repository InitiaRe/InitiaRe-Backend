package middleware

import (
	"github.com/Ho-Minh/InitiaRe-website/config"
	"github.com/Ho-Minh/InitiaRe-website/internal/auth"
)

// Middleware manager
type MiddlewareManager struct {
	cfg    *config.Config
	authUC auth.UseCase
}

// Middleware manager constructor
func NewMiddlewareManager(cfg *config.Config, authUC auth.UseCase) *MiddlewareManager {
	return &MiddlewareManager{
		cfg:    cfg,
		authUC: authUC,
	}
}
