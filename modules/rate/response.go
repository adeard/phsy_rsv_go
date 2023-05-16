package rate

type RateResponse struct {
	ID     int `json:"id"`
	UserId int `json:"user_id"`
	Rates  int `json:"rates"`
}
