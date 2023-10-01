package main

import (
	"fmt"
	"os"

	"github.com/Ho-Minh/InitiaRe-website/config"
	article "github.com/Ho-Minh/InitiaRe-website/internal/article/entity"
	articleCategory "github.com/Ho-Minh/InitiaRe-website/internal/article_category/entity"
	user "github.com/Ho-Minh/InitiaRe-website/internal/auth/entity"
	category "github.com/Ho-Minh/InitiaRe-website/internal/category/entity"
	"github.com/Ho-Minh/InitiaRe-website/internal/models"
	storage "github.com/Ho-Minh/InitiaRe-website/internal/storage/entity"
	todo "github.com/Ho-Minh/InitiaRe-website/internal/todo/entity"
	userInfo "github.com/Ho-Minh/InitiaRe-website/internal/user_info/entity"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type migrateCfg struct {
	seedData            bool
	migrateSchema       bool
	migrateRelationship bool
}

func main() {

	// Init log
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	cfg := migrateCfg{
		seedData:            true,
		migrateSchema:       true,
		migrateRelationship: true,
	}

	// Connect to database
	c := config.GetConfig()
	sslrootcert := "./cert/db_pub_cert.pem"
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=verify-full sslrootcert=%s", c.PostgreSQL.Host, c.PostgreSQL.User, c.PostgreSQL.Password, c.PostgreSQL.DBName, c.PostgreSQL.Port, sslrootcert)

	log.Info().Msg("Connect to database ...")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		log.Fatal().Msg("failed to connect database")
	}

	log.Info().Msgf("Connect to database [%v] successfully", c.PostgreSQL.DBName)

	// Migrate the schema
	if cfg.migrateSchema {
		log.Info().Msg("Run migrate schema ...")

		if !db.Migrator().HasTable(&user.User{}) {
			log.Info().Msg("Create table user ...")
			db.Migrator().CreateTable(&user.User{})
		}

		if !db.Migrator().HasTable(&todo.Todo{}) {
			log.Info().Msg("Create table todo ...")
			db.Migrator().CreateTable(&todo.Todo{})
		}

		if !db.Migrator().HasTable(&userInfo.UserInfo{}) {
			log.Info().Msg("Create table user_info ...")
			db.Migrator().CreateTable(&userInfo.UserInfo{})
		}

		if !db.Migrator().HasTable(&article.Article{}) {
			log.Info().Msg("Create table article ...")
			db.Migrator().CreateTable(&article.Article{})
		}

		if !db.Migrator().HasTable(&articleCategory.ArticleCategory{}) {
			log.Info().Msg("Create table article_category ...")
			db.Migrator().CreateTable(&articleCategory.ArticleCategory{})
		}

		if !db.Migrator().HasTable(&models.Rating{}) {
			log.Info().Msg("Create table rating ...")
			db.Migrator().CreateTable(&models.Rating{})
		}

		if !db.Migrator().HasTable(&category.Category{}) {
			log.Info().Msg("Create table category ...")
			db.Migrator().CreateTable(&category.Category{})
		}

		if !db.Migrator().HasTable(&storage.Storage{}) {
			log.Info().Msg("Create table storage ...")
			db.Migrator().CreateTable(&storage.Storage{})
		}

		if !db.Migrator().HasTable(&models.Status{}) {
			log.Info().Msg("Create table status ...")
			db.Migrator().CreateTable(&models.Status{})
		}

		log.Info().Msg("Migrate successfully")
	}

	// Migrate relationship
	if cfg.migrateRelationship {
		log.Info().Msg("Run migrate relationship ...")

		if db.Migrator().HasTable(&article.Article{}) && db.Migrator().HasTable(&category.Category{}) {
			if !db.Migrator().HasConstraint(&article.Article{}, "initiare_article_category_id_fk") {
				log.Info().Msg("Create constraint table article (category-id-fk) ...")
				db.Exec("ALTER TABLE public.\"initiaRe_article\" ADD CONSTRAINT initiare_article_category_id_fk FOREIGN KEY (category_id) REFERENCES public.\"initiaRe_category\"(id) ON DELETE SET NULL")
			}
		}

		if db.Migrator().HasTable(&article.Article{}) && db.Migrator().HasTable(&category.Category{}) && db.Migrator().HasTable(&articleCategory.ArticleCategory{}) {
			if !db.Migrator().HasConstraint(&articleCategory.ArticleCategory{}, "initiare_article_id_fk") {
				log.Info().Msg("Create constraint table article_category (article-id-fk) ...")
				db.Exec("ALTER TABLE public.\"initiaRe_article_category\" ADD CONSTRAINT initiare_article_id_fk FOREIGN KEY (article_id) REFERENCES public.\"initiaRe_article\"(id) ON DELETE CASCADE")
			}

			if !db.Migrator().HasConstraint(&articleCategory.ArticleCategory{}, "initiare_category_id_fk") {
				log.Info().Msg("Create constraint table article_category (category-id-fk) ...")
				db.Exec("ALTER TABLE public.\"initiaRe_article_category\" ADD CONSTRAINT initiare_category_id_fk FOREIGN KEY (category_id) REFERENCES public.\"initiaRe_category\"(id) ON DELETE CASCADE")
			}
		}

		if db.Migrator().HasTable(&models.Rating{}) && db.Migrator().HasTable(&article.Article{}) && db.Migrator().HasTable(&user.User{}) {
			if !db.Migrator().HasConstraint(&models.Rating{}, "initiare_user_id_fk") {
				log.Info().Msg("Create constraint table rating (user-id-fk) ...")
				db.Exec("ALTER TABLE public.\"initiaRe_rating\" ADD CONSTRAINT initiare_user_id_fk FOREIGN KEY (user_id) REFERENCES public.\"initiaRe_user\"(id) ON DELETE CASCADE")
			}

			if !db.Migrator().HasConstraint(&models.Rating{}, "initiare_article_id_fk") {
				log.Info().Msg("Create constraint table rating (article-id-fk) ...")
				db.Exec("ALTER TABLE public.\"initiaRe_rating\" ADD CONSTRAINT initiare_article_id_fk FOREIGN KEY (article_id) REFERENCES public.\"initiaRe_article\"(id) ON DELETE CASCADE")
			}
		}

		if db.Migrator().HasTable(&userInfo.UserInfo{}) && db.Migrator().HasTable(&user.User{}) {
			if !db.Migrator().HasConstraint(&userInfo.UserInfo{}, "initiare_user_id_fk") {
				log.Info().Msg("Create constraint table user_info (user-id-fk) ...")
				db.Exec("ALTER TABLE public.\"initiaRe_user_info\" ADD CONSTRAINT initiare_user_id_fk FOREIGN KEY (user_id) REFERENCES public.\"initiaRe_user\"(id) ON DELETE CASCADE")
			}
		}

		log.Info().Msg("Migrate successfully")
	}

	// Seeding data
	if cfg.seedData {
		log.Info().Msg("Seeding data ...")
		seedingCategory(db)
		seedingStatus(db)
		log.Info().Msg("Seeding successfully")
	}
}

func seedingCategory(db *gorm.DB) {
	objs := []*category.Category{
		{Id: 1, CategoryName: "Mathematics"},
		{Id: 2, CategoryName: "English Literature"},
		{Id: 3, CategoryName: "Vietnamese Literature"},
		{Id: 4, CategoryName: "Physics"},
		{Id: 5, CategoryName: "Chemistry"},
		{Id: 6, CategoryName: "Biology"},
		{Id: 7, CategoryName: "Social Study"},
		{Id: 8, CategoryName: "History"},
		{Id: 9, CategoryName: "IR & Politics"},
	}
	result := db.Create(&objs)
	if result.Error != nil {
		log.Fatal().Msg(result.Error.Error())
	}
}

func seedingStatus(db *gorm.DB) {
	objs := []*models.Status{
		{StatusId: 1, Category: "article", StatusName: "Pending"},
		{StatusId: 2, Category: "article", StatusName: "Approved"},
		{StatusId: 3, Category: "article", StatusName: "Hidden"},
		{StatusId: 1, Category: "user", StatusName: "Active"},
		{StatusId: 2, Category: "user", StatusName: "Inactive"},
	}
	result := db.Create(&objs)
	if result.Error != nil {
		log.Fatal().Msg(result.Error.Error())
	}
}
