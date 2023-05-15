package user

type UserResponse struct {
	ID          int    `json:"id"`
	Username    string `json:"username"`
	UserTypeId  int    `json:"user_type_id"`
	UserLevelId int    `json:"user_level_id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Address     string `json:"address"`
	ImgProfile  string `json:"img_profile"`
	Gender      string `json:"gender"`
	BirthDate   string `json:"birth_date"`
	IsActive    bool   `json:"is_active"`
}
