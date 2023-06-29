package main

import (
	"fmt"
	"time"

	"github.com/Ho-Minh/InitiaRe-website/config"
	"github.com/Ho-Minh/InitiaRe-website/internal/server"

	"github.com/labstack/gommon/log"
	"github.com/redis/go-redis/v9"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//	@title			Swagger Clean Architecture API
//	@version		1.0
//	@description	Example Golang REST API.

//	@contact.name	Vu Ky
//	@contact.url	https://github.com/vukyn
//	@contact.email	vukynpro@gmailcom

// @BasePath	api/v1
func main() {
	log.Info("Starting api server")

	cfg := config.GetConfig()

	// Init Postgresql
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", cfg.PostgreSQL.Host, cfg.PostgreSQL.User, cfg.PostgreSQL.Password, cfg.PostgreSQL.DBName, cfg.PostgreSQL.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Postgresql init: %s", err)
	}

	// Init Redis
	redisClient := redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port),
		MinIdleConns: cfg.Redis.MinIdleConns,
		PoolSize:     cfg.Redis.PoolSize,
		PoolTimeout:  time.Duration(cfg.Redis.PoolTimeout) * time.Second,
	})
	defer redisClient.Close()

	s := server.NewServer(cfg, db, redisClient)
	if err = s.Run(); err != nil {
		log.Fatal(err)
	}
}
