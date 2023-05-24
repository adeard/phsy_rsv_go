package registry

import (
	"phsy_rsv_go/modules/city"

	"gorm.io/gorm"
)

func CityRegistry(db *gorm.DB) city.Service {
	cityRepository := city.NewRepository(db)
	cityService := city.NewService(cityRepository)

	return cityService
}
