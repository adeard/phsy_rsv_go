package seeder

import (
	"phsy_rsv_go/domain"

	"gorm.io/gorm"
)

func InsertUserType(Db *gorm.DB) {
	var total int64
	var ut domain.UserType
	Db.Model(&ut).Count(&total)
	if total == 0 {
		var data = []domain.UserType{
			{
				Name: "patient", IsActive: true,
			},
			{
				Name: "staff", IsActive: true,
			},
		}

		Db.Create(&data)
	}
}
