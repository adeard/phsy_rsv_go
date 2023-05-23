package user

import (
	userlevel "phsy_rsv_go/modules/user_level"
	usertype "phsy_rsv_go/modules/user_type"
)

type UserResponse struct {
	ID          int                         `json:"id"`
	Username    string                      `json:"username"`
	UserTypeId  uint                        `json:"user_type_id"`
	UserLevelId int                         `json:"user_level_id"`
	Name        string                      `json:"name"`
	Email       string                      `json:"email"`
	Address     string                      `json:"address"`
	ImgProfile  string                      `json:"img_profile"`
	Gender      string                      `json:"gender"`
	BirthDate   string                      `json:"birth_date"`
	IsActive    bool                        `json:"is_active"`
	UserType    usertype.UserTypeResponse   `json:"user_type"`
	UserLevel   userlevel.UserLevelResponse `json:"user_level"`
}
