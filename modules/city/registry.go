package city

import (
	"gorm.io/gorm"
)

func CityRegistry(db *gorm.DB) Service {
	cityRepository := NewRepository(db)
	cityService := NewService(cityRepository)

	return cityService
}
