package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Ho-Minh/InitiaRe-website/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

const (
	ctxTimeout = 5
)

// Server struct
type Server struct {
	echo        *echo.Echo
	cfg         *config.Config
	db          *gorm.DB
	redisClient *redis.Client
}

// constructor
func NewServer(cfg *config.Config, db *gorm.DB, redisClient *redis.Client) *Server {
	return &Server{echo: echo.New(), cfg: cfg, db: db, redisClient: redisClient}
}

func (s *Server) Run() error {

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", s.cfg.Server.Port),
		ReadTimeout:  time.Second * s.cfg.Server.ReadTimeout,
		WriteTimeout: time.Second * s.cfg.Server.WriteTimeout,
	}

	go func() {
		log.Infof("Server is listening on PORT: %d", s.cfg.Server.Port)
		if err := s.echo.StartServer(server); err != nil {
			log.Fatalf("Error starting Server: ", err)
		}
	}()

	if err := s.MapHandlers(s.echo); err != nil {
		return err
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), ctxTimeout*time.Second)
	defer shutdown()

	log.Info("Server Exited Properly")
	return s.echo.Server.Shutdown(ctx)
}
