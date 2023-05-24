package registry

import (
	"phsy_rsv_go/modules/province"

	"gorm.io/gorm"
)

func ProvinceRegistry(db *gorm.DB) province.Service {
	provinceRepository := province.NewRepository(db)
	provinceService := province.NewService(provinceRepository)

	return provinceService
}
