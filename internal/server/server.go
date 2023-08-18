package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Ho-Minh/InitiaRe-website/config"

	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
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
	ctn         *azblob.Client
	redisClient *redis.Client
}

func NewServer(cfg *config.Config, db *gorm.DB, ctn *azblob.Client, redisClient *redis.Client) *Server {
	return &Server{
		echo:        echo.New(),
		cfg:         cfg,
		db:          db,
		ctn:         ctn,
		redisClient: redisClient,
	}
}

func (s *Server) Run() error {

	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", s.cfg.Server.Port),
		ReadTimeout:  time.Second * time.Duration(s.cfg.Server.ReadTimeout),
		WriteTimeout: time.Second * time.Duration(s.cfg.Server.WriteTimeout),
	}

	go func() {
		log.Info().Msgf("Server is listening on PORT: %v", s.cfg.Server.Port)
		if err := s.echo.StartServer(server); err != nil {
			log.Fatal().Msgf("Error starting server: %v", err)
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

	log.Info().Msg("Server Exited Properly")
	return s.echo.Server.Shutdown(ctx)
}
