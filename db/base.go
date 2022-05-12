package db

import (
	"r-cha/goblog/models"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var database *gorm.DB
	var err error

	// settings := config.GetSettings()
	env := "local" // settings.environment
	if env == "prod" {
		dsn := "connectionstring"
		database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	} else if env == "local" {
		database, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	}

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&models.Post{})

	DB = database
}
