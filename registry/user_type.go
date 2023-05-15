package registry

import (
	usertype "phsy_rsv_go/modules/user_type"

	"gorm.io/gorm"
)

func UserTypeRegistry(db *gorm.DB) usertype.Service {
	userTypeRepository := usertype.NewRepository(db)
	userTypeService := usertype.NewService(userTypeRepository)

	return userTypeService
}
