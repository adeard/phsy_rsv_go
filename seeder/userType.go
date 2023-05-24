package seeder

import (
	usertype "phsy_rsv_go/modules/user_type"

	"gorm.io/gorm"
)

func InsertUserType(Db *gorm.DB) {
	var total int64
	var ut usertype.UserType
	Db.Model(&ut).Count(&total)
	if total == 0 {
		var data = []usertype.UserType{
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
