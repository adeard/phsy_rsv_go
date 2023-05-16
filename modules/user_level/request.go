package userlevel

type UserLevelRequest struct {
	Name     string `gorm:"size:255;not null" json:"name" binding:"required"`
	IsActive bool   `gorm:"default:true;" json:"is_active"`
}
