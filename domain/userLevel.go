package domain

import "gorm.io/gorm"

type UserLevel struct {
	gorm.Model
	Name     string `gorm:"size:255;not null" json:"name"`
	IsActive bool   `gorm:"default:true;" json:"is_active"`
}

type UserLevelRequest struct {
	Name     string `gorm:"size:255;not null" json:"name" binding:"required"`
	IsActive bool   `gorm:"default:true;" json:"is_active"`
}

type UserLevelResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
}
