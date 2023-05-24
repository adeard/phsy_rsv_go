package domain

import "gorm.io/gorm"

type Rate struct {
	gorm.Model
	UserId int `gorm:"not null" json:"user_id"`
	Rates  int `gorm:"not null" json:"rates"`
}

type RateRequest struct {
	UserId int `gorm:"not null" json:"user_id" binding:"required"`
	Rates  int `gorm:"not null" json:"rates"  binding:"required"`
}

type RateResponse struct {
	ID     int `json:"id"`
	UserId int `json:"user_id"`
	Rates  int `json:"rates"`
}
