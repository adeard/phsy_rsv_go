package city

type CityResponse struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	IsActive   bool   `json:"is_active"`
	ProvinceId int    `json:"province_id"`
}
