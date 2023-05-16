package registry

import (
	"phsy_rsv_go/modules/rate"

	"gorm.io/gorm"
)

func RateRegistry(db *gorm.DB) rate.Service {
	rateRepository := rate.NewRepository(db)
	rateService := rate.NewService(rateRepository)

	return rateService
}
