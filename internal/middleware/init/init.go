package init

import (
	"github.com/Ho-Minh/InitiaRe-website/config"
	initAuth "github.com/Ho-Minh/InitiaRe-website/internal/auth/init"
	"github.com/Ho-Minh/InitiaRe-website/internal/middleware"
)

type Init struct {
	MiddlewareManager middleware.IMiddlewareManager
}

func NewInit(
	cfg *config.Config,
	initAuth *initAuth.Init,
) *Init {
	mw := middleware.NewMiddlewareManager(cfg, initAuth.CacheRepository)
	return &Init{
		MiddlewareManager: mw,
	}
}
