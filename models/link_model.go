package model

import (
	"time"

)

type Shortsy struct {
	Id          uint   `gorm:"primaryKey" json:"id"`
	OriginalURL string `json:"original_url"`
	ShortCode   string `json:"short_code" gorm:"uniqueIndex"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
