package seeder

import (
	"phsy_rsv_go/modules/city"

	"gorm.io/gorm"
)

func InsertCity(Db *gorm.DB) {
	var total int64
	var ul city.City
	Db.Model(&ul).Count(&total)
	if total == 0 {
		var data = []city.City{
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
