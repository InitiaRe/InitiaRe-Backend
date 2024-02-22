package init

import (
	"InitiaRe-website/config"
	initAuth "InitiaRe-website/internal/auth/init"
	"InitiaRe-website/internal/middleware"
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
