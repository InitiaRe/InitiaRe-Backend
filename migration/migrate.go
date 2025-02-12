package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"InitiaRe-website/config"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	article "InitiaRe-website/internal/article/entity"
	articleCategory "InitiaRe-website/internal/article_category/entity"
	user "InitiaRe-website/internal/auth/entity"
	category "InitiaRe-website/internal/category/entity"
	"InitiaRe-website/internal/models"
	rating "InitiaRe-website/internal/rating/entity"
	school "InitiaRe-website/internal/school/entity"
	storage "InitiaRe-website/internal/storage/entity"
	todo "InitiaRe-website/internal/todo/entity"
	userInfo "InitiaRe-website/internal/user_info/entity"
)

func main() {
	cfg := config.GetConfig()

	// Init Logger
	initLogger(&cfg.Logger)

	// Init DB
	db := initDB(&cfg.PostgreSQL)

	// Start migration
	migrate(db)
}

func migrate(db *gorm.DB) {
	log.Info().Msg("Start migration")

	if db.Migrator().HasTable(&user.User{}) {
		log.Info().Msg("User table already exists. Skipping migration.")
	} else {
		log.Info().Msg("User table does not exist. Running migration.")
		db.AutoMigrate(&user.User{})
	}

	if db.Migrator().HasTable(&article.Article{}) {
		log.Info().Msg("Article table already exists. Skipping migration.")
	} else {
		log.Info().Msg("Article table does not exist. Running migration.")
		db.AutoMigrate(&article.Article{})
	}

	if db.Migrator().HasTable(&user.User{}) {
		log.Info().Msg("User table already exists. Skipping migration.")
	} else {
		log.Info().Msg("User table does not exist. Running migration.")
		db.AutoMigrate(&user.User{})
	}

	if db.Migrator().HasTable(&articleCategory.ArticleCategory{}) {
		log.Info().Msg("ArticleCategory table already exists. Skipping migration.")
	} else {
		log.Info().Msg("ArticleCategory table does not exist. Running migration.")
		db.AutoMigrate(&articleCategory.ArticleCategory{})
	}

	if db.Migrator().HasTable(&category.Category{}) {
		log.Info().Msg("Category table already exists. Skipping migration.")
	} else {
		log.Info().Msg("Category table does not exist. Running migration.")
		db.AutoMigrate(&category.Category{})
	}

	if db.Migrator().HasTable(&models.Status{}) {
		log.Info().Msg("Status table already exists. Skipping migration.")
	} else {
		log.Info().Msg("Status table does not exist. Running migration.")
		db.AutoMigrate(&models.Status{})
	}

	if db.Migrator().HasTable(&storage.Storage{}) {
		log.Info().Msg("Storage table already exists. Skipping migration.")
	} else {
		log.Info().Msg("Storage table does not exist. Running migration.")
		db.AutoMigrate(&storage.Storage{})
	}

	if db.Migrator().HasTable(&todo.Todo{}) {
		log.Info().Msg("Todo table already exists. Skipping migration.")
	} else {
		log.Info().Msg("Todo table does not exist. Running migration.")
		db.AutoMigrate(&todo.Todo{})
	}

	if db.Migrator().HasTable(&school.School{}) {
		log.Info().Msg("School table already exists. Skipping migration.")
	} else {
		log.Info().Msg("School table does not exist. Running migration.")
		db.AutoMigrate(&school.School{})
	}

	if db.Migrator().HasTable(&rating.Rating{}) {
		log.Info().Msg("Rating table already exists. Skipping migration.")
	} else {
		log.Info().Msg("Rating table does not exist. Running migration.")
		db.AutoMigrate(&rating.Rating{})
	}

	if db.Migrator().HasTable(&userInfo.UserInfo{}) {
		log.Info().Msg("UserInfo table already exists. Skipping migration.")
	} else {
		log.Info().Msg("UserInfo table does not exist. Running migration.")
		db.AutoMigrate(&userInfo.UserInfo{})
	}

	log.Info().Msg("Auto migrate completed!")
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
