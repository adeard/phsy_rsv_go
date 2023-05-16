package userlevel

type UserLevelResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
}
