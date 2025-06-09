package service

import (	
	model "shortsy/models"
	"shortsy/database"
	"github.com/matoous/go-nanoid/v2"
)

func PostLink(original_url string)(*model.Shortsy,error){

	id,_ := gonanoid.New()
	short := id[:6]

	newUrl := &model.Shortsy{
		OriginalURL:original_url,
		ShortCode: short,
	}
	if err := database.DB.Create(newUrl).Error; err != nil {
		return nil, err
	}
	return newUrl,nil
}

func GetLink(shortCode string) (*model.Shortsy,error){
	var shortLink model.Shortsy
	result := database.DB.Where("short_code = ?",shortCode).First(&shortLink)

	if result.Error != nil {
		return nil,result.Error
	}
	return &shortLink,nil
}