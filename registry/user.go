package registry

import (
	"phsy_rsv_go/modules/user"

	"gorm.io/gorm"
)

func UserRegistry(db *gorm.DB) user.Service {
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	return userService
}
