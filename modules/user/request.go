package user

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
