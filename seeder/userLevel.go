package seeder

import (
	userlevel "phsy_rsv_go/modules/user_level"

	"gorm.io/gorm"
)

func InsertUserLevel(Db *gorm.DB) {
	var total int64
	var ul userlevel.UserLevel
	Db.Model(&ul).Count(&total)
	if total == 0 {
		var data = []userlevel.UserLevel{
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
