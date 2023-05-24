package province

import "gorm.io/gorm"

func ProvinceRegistry(db *gorm.DB) Service {
	provinceRepository := NewRepository(db)
	provinceService := NewService(provinceRepository)

	return provinceService
}
