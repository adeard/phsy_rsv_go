package seeder

import (
	"phsy_rsv_go/domain"

	"gorm.io/gorm"
)

func InsertCity(Db *gorm.DB) {
	var total int64
	var ul domain.City
	Db.Model(&ul).Count(&total)
	if total == 0 {
		var data = []domain.City{
			{
				Name: "Jakarta Timur", IsActive: true, ProvinceId: 17,
			},
			{
				Name: "Jakarta Selatan", IsActive: true, ProvinceId: 17,
			},
		}

		Db.Create(&data)
	}
}
