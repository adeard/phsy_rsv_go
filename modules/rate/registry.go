package rate

import "gorm.io/gorm"

func RateRegistry(db *gorm.DB) Service {
	rateRepository := NewRepository(db)
	rateService := NewService(rateRepository)

	return rateService
}
