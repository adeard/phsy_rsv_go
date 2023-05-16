package rate

import "gorm.io/gorm"

type Rate struct {
	gorm.Model
	UserId int `gorm:"not null" json:"user_id"`
	Rates  int `gorm:"not null" json:"rates"`
}
