package database

import (
	"log"
	"os"
	model "shortsy/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := os.Getenv("DB_URL") 
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}

	DB = db
	log.Println("✅ Connected to PostgreSQL with GORM")

	db.AutoMigrate(&model.Shortsy{})
}
