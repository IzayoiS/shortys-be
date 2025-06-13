package database

import (
	"errors"
	"log"

	model "shortsy/models"

	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {
	if db == nil {
		log.Fatal("❌ Database connection is nil (db == nil)")
	}

	shorts := []model.Shortsy{
		{OriginalURL: "https://google.com", ShortCode: "ggl"},
		{OriginalURL: "https://github.com", ShortCode: "ghub"},
	}

	for _, s := range shorts {
		var existing model.Shortsy
		err := db.Where("short_code = ?", s.ShortCode).First(&existing).Error

		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err := db.Create(&s).Error; err != nil {
				log.Printf("❌ Failed to seed: %+v\n", err)
			} else {
				log.Printf("✅ Seeded: %s → %s\n", s.ShortCode, s.OriginalURL)
			}
		} else if err != nil {
			log.Printf("⚠️ Error checking existing record: %+v\n", err)
		} else {
			log.Printf("ℹ️ Skipped existing: %s\n", s.ShortCode)
		}
	}
}
