package models

import "gorm.io/gorm"

type Shortsy struct {
	gorm.Model
	OriginalURL string `json:"original_url"`
	ShortCode string `json:"short_code" gorm:"uniqueIndex"`
}

