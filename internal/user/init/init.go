package init

import (
	"github.com/Ho-Minh/InitiaRe-website/config"
	initMW "github.com/Ho-Minh/InitiaRe-website/internal/middleware/init"
	handler "github.com/Ho-Minh/InitiaRe-website/internal/user/delivery/http"

	"gorm.io/gorm"
)

type Init struct {
	Handler handler.IHandler
}

func NewInit(
	db *gorm.DB,
	cfg *config.Config,
	mw *initMW.Init,
) *Init {
	handler := handler.InitHandler(cfg, mw.MiddlewareManager)
	return &Init{
		Handler: handler,
	}
}
