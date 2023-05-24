package province

import "gorm.io/gorm"

type Province struct {
	gorm.Model
	Name     string `gorm:"size:255" json:"name"`
	IsActive bool   `gorm:"default:1" json:"is_active"`
}
