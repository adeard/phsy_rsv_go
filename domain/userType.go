package domain

import "gorm.io/gorm"

type UserType struct {
	gorm.Model
	Name     string `gorm:"size:255;not null" json:"name"`
	IsActive bool   `gorm:"default:true;" json:"is_active"`
}

type UserTypeRequest struct {
	Name     string `gorm:"size:255;not null" json:"name" binding:"required"`
	IsActive bool   `gorm:"default:true;" json:"is_active" binding:"required"`
}

type UserTypeResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
}
