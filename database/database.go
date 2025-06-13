package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	model "shortsy/models"
)

var DB *gorm.DB

func Connect() {
	
	db, err := gorm.Open(postgres.Open("postgresql://postgres:3RVwLUISRcri25w6@db.jgsyvsxijotyuoxxogsa.supabase.co:5432/postgres"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database")
	}

	DB = db

	db.AutoMigrate(&model.Shortsy{})
}
