package userlevel

import "gorm.io/gorm"

func UserLevelRegistry(db *gorm.DB) Service {
	userLevelRepository := NewRepository(db)
	userLevelService := NewService(userLevelRepository)

	return userLevelService
}
