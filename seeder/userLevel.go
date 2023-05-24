package seeder

import (
	"phsy_rsv_go/domain"

	"gorm.io/gorm"
)

func InsertUserLevel(Db *gorm.DB) {
	var total int64
	var ul domain.UserLevel
	Db.Model(&ul).Count(&total)
	if total == 0 {
		var data = []domain.UserLevel{
			{
				Name: "admin", IsActive: true,
			},
			{
				Name: "user", IsActive: true,
			},
		}

		Db.Create(&data)
	}
}
