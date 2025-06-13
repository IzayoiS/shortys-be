package database

import (
	"log"
	model "shortsy/models"

	"gorm.io/gorm"
)

func Seed(db*gorm.DB){
    if DB == nil {
		log.Fatal("Failed to connect to database (DB is nil)")
	}
    
	shorts := []model.Shortsy{
		{OriginalURL: "https://google.com", ShortCode: "ggl"},
		{OriginalURL: "https://github.com", ShortCode: "ghub"},
	}
	 for _, s := range shorts {
       
        var existing model.Shortsy
        err := db.Where("short_code = ?", s.ShortCode).First(&existing).Error
        if err == gorm.ErrRecordNotFound {
       
            err := db.Create(&s).Error
            if err != nil {
                log.Println("Seeding fail:", err)
            }
        }
    }
}
