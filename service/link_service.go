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