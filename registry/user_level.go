package registry

import (
	userlevel "phsy_rsv_go/modules/user_level"

	"gorm.io/gorm"
)

func UserLevelRegistry(db *gorm.DB) userlevel.Service {
	userLevelRepository := userlevel.NewRepository(db)
	userLevelService := userlevel.NewService(userLevelRepository)

	return userLevelService
}
