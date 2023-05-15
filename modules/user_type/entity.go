package usertype

import (
	"gorm.io/gorm"
)

type UserType struct {
	gorm.Model
	Name     string `gorm:"size:255;not null" json:"name"`
	IsActive bool   `gorm:"default:true;" json:"is_active"`
}
