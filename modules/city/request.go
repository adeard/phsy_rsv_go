package city

type CityRequest struct {
	Name       string `json:"name" binding:"required"`
	ProvinceId int    `json:"province_id" binding:"required"`
	IsActive   bool   `gorm:"default:true" json:"is_active"`
}
