package userlevel

import (
	"gorm.io/gorm"
)

type UserLevel struct {
	gorm.Model
	Name     string `gorm:"size:255;not null" json:"name"`
	IsActive bool   `gorm:"default:true;" json:"is_active"`
}
