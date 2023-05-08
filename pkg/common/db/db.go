package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"pizza/pkg/common/models"
)

// Init Start the application
func Init(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Pizza{})

	return db
}
