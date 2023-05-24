package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username    string    `gorm:"size:255;not null;unique" json:"username"`
	Password    string    `gorm:"size:255;not null;" json:"password"`
	UserTypeId  uint      `gorm:"default:1" json:"user_type_id"`
	UserLevelId int       `gorm:"default:1" json:"user_level_id"`
	Name        string    `gorm:"size:255;default:null" json:"name"`
	Email       string    `gorm:"size:255;default:null" json:"email"`
	Address     string    `gorm:"default:null" json:"address"`
	ImgProfile  string    `gorm:"size:255;" json:"img_profile"`
	Gender      string    `gorm:"size:255;default:male" json:"gender"`
	BirthDate   string    `gorm:"type:date;default:null" json:"birth_date"`
	IsActive    bool      `gorm:"default:true;" json:"is_active"`
	UserType    UserType  `gorm:"foreignkey:UserTypeId;references:ID" json:"user_type"`
	UserLevel   UserLevel `gorm:"foreignkey:UserLevelId;references:ID" json:"user_level"`
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateRequest struct {
	Username    string `json:"username"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	UserTypeId  int    `json:"user_type_id"`
	UserLevelId int    `json:"user_level_id"`
	Address     string `json:"address"`
	ImgProfile  string `json:"img_profile"`
	Gender      string `json:"gender"`
	BirthDate   string `json:"birth_date"`
	IsActive    bool   `json:"is_active"`
}

type UserResponse struct {
	ID          int               `json:"id"`
	Username    string            `json:"username"`
	UserTypeId  uint              `json:"user_type_id"`
	UserLevelId int               `json:"user_level_id"`
	Name        string            `json:"name"`
	Email       string            `json:"email"`
	Address     string            `json:"address"`
	ImgProfile  string            `json:"img_profile"`
	Gender      string            `json:"gender"`
	BirthDate   string            `json:"birth_date"`
	IsActive    bool              `json:"is_active"`
	UserType    UserTypeResponse  `json:"user_type"`
	UserLevel   UserLevelResponse `json:"user_level"`
}
