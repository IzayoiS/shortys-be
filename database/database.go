package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	model "shortsy/models"
)

var DB *gorm.DB

func Connect() {
	
	db, err := gorm.Open(postgres.Open("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6InRwcHZ6c2pzdndibWRkZXd0bHh2Iiwicm9sZSI6ImFub24iLCJpYXQiOjE3NDkxMzU3MjksImV4cCI6MjA2NDcxMTcyOX0.VnMrcMNgHCZs8uCp8FHnvq0RQAi4HZqn-E4mjd_oHhI"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database")
	}

	DB = db

	db.AutoMigrate(&model.Shortsy{})
}
