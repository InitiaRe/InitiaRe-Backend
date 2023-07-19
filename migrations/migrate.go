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
)

func main() {

	// Connect to database
	c := config.GetConfig()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", c.PostgreSQL.Host, c.PostgreSQL.User, c.PostgreSQL.Password, c.PostgreSQL.DBName, c.PostgreSQL.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
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
		{Id: 1, Category: "Mathematics"},
		{Id: 2, Category: "English Literature"},
		{Id: 3, Category: "Vietnamese Literature"},
		{Id: 4, Category: "Physics"},
		{Id: 5, Category: "Chemistry"},
		{Id: 6, Category: "Biology"},
		{Id: 7, Category: "Social Study"},
		{Id: 8, Category: "History"},
		{Id: 9, Category: "IR & Politics"},
	}
	result := db.Create(&objs)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
}

func seedingStatus(db *gorm.DB) {
	objs := []*models.Status{
		{Id: 1, Category: "article", StatusName: "Pending"},
		{Id: 2, Category: "article", StatusName: "Approved"},
		{Id: 3, Category: "article", StatusName: "Hidden"},
		{Id: 1, Category: "user", StatusName: "Active"},
		{Id: 2, Category: "user", StatusName: "Inactive"},
	}
	result := db.Create(&objs)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
}
