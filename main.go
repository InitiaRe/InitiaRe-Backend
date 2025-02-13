package main

import (
	"crypto/tls"
	"fmt"
	"os"
	"strconv"
	"time"

	"InitiaRe-website/config"
	"InitiaRe-website/internal/server"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

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

// @BasePath	/api/v1
func main() {
	cfg := config.GetConfig()

	// Init Logger
	initLogger(&cfg.Logger)

	// Init DB
	db := initDB(&cfg.PostgreSQL)

	// Init Redis
	redisClient := initRedis(&cfg.Redis)
	defer redisClient.Close()

	// Init Azure Blob Storage
	ctnClient := initContainer(&cfg.Storage)

	// Init server
	log.Info().Msg("Starting api server")
	s := server.NewServer(cfg, db, ctnClient, redisClient)
	if err := s.Run(); err != nil {
		log.Fatal().Msg(err.Error())
	}
}

func initDB(cfg *config.PostgreSQLConfig) *gorm.DB {
	log.Info().Msg("Init DB")
	zLogger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	zLogger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	newLogger := logger.New(
		&zLogger, // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,         // Don't include params in the SQL log
			Colorful:                  true,          // Disable color
		},
	)
	sslrootcert := "./cert/db_pub_cert.pem"
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=verify-full sslrootcert=%s", cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port, sslrootcert)

	if !cfg.SSL {
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d", cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatal().Msgf("DB init: %s", err)
	}
	return db
}

func initRedis(cfg *config.RedisConfig) *redis.Client {
	log.Info().Msg("Init Redis")

	redisClient := redis.NewClient(&redis.Options{
		Addr:        fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Username:    cfg.Username,
		Password:    cfg.Password,
		DB:          cfg.DB,
		PoolSize:    cfg.PoolSize,
		PoolTimeout: time.Duration(cfg.PoolTimeout) * time.Second,
	})

	if cfg.TLS {
		redisClient = redis.NewClient(&redis.Options{
			Addr:        fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
			Username:    cfg.Username,
			Password:    cfg.Password,
			DB:          cfg.DB,
			PoolSize:    cfg.PoolSize,
			PoolTimeout: time.Duration(cfg.PoolTimeout) * time.Second,
			TLSConfig:   &tls.Config{MinVersion: tls.VersionTLS12},
		})
	}

	return redisClient
}

func initLogger(cfg *config.LoggerConfig) {
	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		short := file
		for i := len(file) - 1; i > 0; i-- {
			if file[i] == '/' {
				short = file[i+1:]
				break
			}
		}
		file = short
		return file + ":" + strconv.Itoa(line)
	}
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.With().Caller().Logger()
	if cfg.Mode == "pretty" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	switch cfg.Level {
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "fatal":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	case "panic":
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
}

func initContainer(cfg *config.StorageConfig) *azblob.Client {
	credential, err := azblob.NewSharedKeyCredential(cfg.AccountName, cfg.AccountKey)
	if err != nil {
		log.Fatal().Msgf("Azure credential init: %s", err)
	}
	client, err := azblob.NewClientWithSharedKeyCredential(cfg.Host, credential, nil)
	if err != nil {
		log.Fatal().Msgf("Azure client init: %s", err)
	}
	return client
}
