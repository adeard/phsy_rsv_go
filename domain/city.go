package domain

import (
	"gorm.io/gorm"
)

type City struct {
	gorm.Model
	Name       string   `gorm:"size:255" json:"name"`
	ProvinceId int      `json:"province_id"`
	IsActive   bool     `gorm:"default:1" json:"is_active"`
	Province   Province `gorm:"foreignkey:ProvinceId;references:ID" json:"province"`
}

type CityRequest struct {
	Name       string `json:"name" binding:"required"`
	ProvinceId int    `json:"province_id" binding:"required"`
	IsActive   bool   `gorm:"default:true" json:"is_active"`
}

type CityResponse struct {
	ID         int              `json:"id"`
	Name       string           `json:"name"`
	IsActive   bool             `json:"is_active"`
	ProvinceId int              `json:"province_id"`
	Province   ProvinceResponse `json:"province"`
}
