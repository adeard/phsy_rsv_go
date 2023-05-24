package domain

import "gorm.io/gorm"

type Province struct {
	gorm.Model
	Name     string `gorm:"size:255" json:"name"`
	IsActive bool   `gorm:"default:true" json:"is_active"`
}

type ProvinceRequest struct {
	Name     string `json:"name" binding:"required"`
	IsActive bool   `gorm:"default:true;" json:"is_active"`
}

type ProvinceResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
}
