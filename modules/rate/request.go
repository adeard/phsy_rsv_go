package rate

type RateRequest struct {
	UserId int `gorm:"not null" json:"user_id" binding:"required"`
	Rates  int `gorm:"not null" json:"rates"  binding:"required"`
}
