package city

import "gorm.io/gorm"

type City struct {
	gorm.Model
	Name       string `gorm:"size:255" json:"name"`
	ProvinceId int    `json:"province_id"`
	IsActive   bool   `gorm:"default:1" json:"is_active"`
}
