package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	model "shortsy/models"
)

var DB *gorm.DB

func Connect() {
	
	dsn := "postgresql://postgres:jTdHHWEMMvl9Cq78@db.brmuhskfoxpfsiegzyyf.supabase.co:5432/postgres"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database")
	}

	DB = db

	db.AutoMigrate(&model.Shortsy{})
}
