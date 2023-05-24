package usertype

import "gorm.io/gorm"

func UserTypeRegistry(db *gorm.DB) Service {
	userTypeRepository := NewRepository(db)
	userTypeService := NewService(userTypeRepository)

	return userTypeService
}
