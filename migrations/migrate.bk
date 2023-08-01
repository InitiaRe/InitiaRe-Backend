package main

import (
	"fmt"

	"github.com/Ho-Minh/InitiaRe-website/config"
	articleEntity "github.com/Ho-Minh/InitiaRe-website/internal/article/entity"
	userEntity "github.com/Ho-Minh/InitiaRe-website/internal/auth/entity"
	categoryEntity "github.com/Ho-Minh/InitiaRe-website/internal/category/entity"
	"github.com/Ho-Minh/InitiaRe-website/internal/models"
	todoEntity "github.com/Ho-Minh/InitiaRe-website/internal/todo/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {

	// Connect to database
	c := config.GetConfig()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", c.PostgreSQL.Host, c.PostgreSQL.User, c.PostgreSQL.Password, c.PostgreSQL.DBName, c.PostgreSQL.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connect to database successfully")

	fmt.Println("Run migrate ...")
	// Migrate the schema
	db.AutoMigrate(&userEntity.User{})
	db.AutoMigrate(&todoEntity.Todo{})
	db.AutoMigrate(&models.UserInfo{})
	db.AutoMigrate(&articleEntity.Article{})
	db.AutoMigrate(&models.Rating{})
	db.AutoMigrate(&categoryEntity.Category{})
	db.AutoMigrate(&models.Status{})
	fmt.Println("Migrate successfully")

	fmt.Println("Seeding data ...")
	// Seeding data
	seedingCategory(db)
	seedingStatus(db)
	fmt.Println("Seeding successfully")
}

func seedingCategory(db *gorm.DB) {
	objs := []*categoryEntity.Category{
		{Category: "Mathematics"},
		{Category: "English Literature"},
		{Category: "Vietnamese Literature"},
		{Category: "Physics"},
		{Category: "Chemistry"},
		{Category: "Biology"},
		{Category: "Social Study"},
		{Category: "History"},
		{Category: "IR & Politics"},
	}
	result := db.Create(&objs)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
}

func seedingStatus(db *gorm.DB) {
	objs := []*models.Status{
		{Category: "article", StatusName: "Pending"},
		{Category: "article", StatusName: "Approved"},
		{Category: "article", StatusName: "Hidden"},
		{Category: "user", StatusName: "Active"},
		{Category: "user", StatusName: "Inactive"},
	}
	result := db.Create(&objs)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
}
