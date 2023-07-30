package main

import (
	"crypto/tls"
	"fmt"
	"os"
	"time"

	"github.com/Ho-Minh/InitiaRe-website/config"
	"github.com/Ho-Minh/InitiaRe-website/internal/server"
	"github.com/joho/godotenv"

	"github.com/labstack/gommon/log"
	"github.com/redis/go-redis/v9"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//	@title			InitiaRe API
//	@version		1.0
//	@description	InitiaRe REST API.

//	@contact.name	Someone here
//	@contact.url	contact.here
//	@contact.email	email@here.com

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization

//	@BasePath	/api/v1
func main() {
	log.Info("Starting api server")

	cfg := config.GetConfig()

	port := os.Getenv("PORT")
	// Get port from .env file in case of running locally
	if port == "" {
		err := godotenv.Load(".env")
		if err != nil {
			panic(err.Error())
		}
		port = os.Getenv("PORT")
		if port == "" {
			log.Fatal("$PORT must be set")
		}
	}
	cfg.Server.Port = port

	// Init Logger
	newLogger := logger.New(
		log.New("GORM:"), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,          // Don't include params in the SQL log
			Colorful:                  false,         // Disable color
		},
	)

	// Init Postgresql
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", cfg.PostgreSQL.Host, cfg.PostgreSQL.User, cfg.PostgreSQL.Password, cfg.PostgreSQL.DBName, cfg.PostgreSQL.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatalf("Postgresql init: %s", err)
	}

	// Init Redis
	redisClient := redis.NewClient(&redis.Options{
		Addr:        fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port),
		Username:    cfg.Redis.Username,
		Password:    cfg.Redis.Password,
		DB:          cfg.Redis.DB,
		PoolSize:    cfg.Redis.PoolSize,
		PoolTimeout: time.Duration(cfg.Redis.PoolTimeout) * time.Second,
		TLSConfig:   &tls.Config{MinVersion: tls.VersionTLS12},
	})
	defer redisClient.Close()

	s := server.NewServer(cfg, db, redisClient)
	if err = s.Run(); err != nil {
		log.Fatal(err)
	}
}
