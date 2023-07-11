package middleware

import (
	"github.com/Ho-Minh/InitiaRe-website/config"
	authRepo "github.com/Ho-Minh/InitiaRe-website/internal/auth/repository"
)

// Middleware manager
type MiddlewareManager struct {
	cfg      *config.Config
	authRepo authRepo.IRepository
}

// Middleware manager constructor
func NewMiddlewareManager(cfg *config.Config, authRepo authRepo.IRepository) *MiddlewareManager {
	return &MiddlewareManager{
		cfg:      cfg,
		authRepo: authRepo,
	}
}