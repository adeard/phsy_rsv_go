package province

type ProvinceRequest struct {
	Name     string `json:"name" binding:"required"`
	IsActive bool   `gorm:"default:true;" json:"is_active"`
}
