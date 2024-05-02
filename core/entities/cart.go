package entities

type CartsAddReq struct {
	UserName      string  `form:"user_name"`
	ShowTimeId    int64   `form:"show_time_id"`
	SeatsPosition string  `form:"seats_position"`
	Price         float64 `form:"price"`
}
type CartsAddResp struct {
	Result Result `json:"result"`
}
type CartFindByFormResp struct {
	Result Result   `json:"result"`
	Carts  []*Carts `json:"carts"`
}
type Carts struct {
	Id             int64   `json:"id"`
	ShowTimeId     int64   `json:"show_time_id"`
	SeatsPosition  string  `json:"seats_position"`
	Price          float64 `json:"price"`
	MovieTime      int     `json:"movie_time"`
	CinemaName     string  `json:"cinema_name"`
	Description    string  `json:"description"`
	Conscious      string  `json:"conscious"`
	District       string  `json:"district"`
	Commune        string  `json:"commune"`
	AddressDetails string  `json:"address_details"`
	MovieName      string  `json:"movie_name"`
	MovieDuration  int     `json:"movie_duration"`
	AgeLimit       int     `json:"age_limit"`
}
type CartsUpdateReq struct {
	Id            int64  `form:"id"`
	UserName      string `form:"user_name"`
	ShowTimeId    int64  `form:"show_time_id"`
	SeatsPosition string `form:"seats_position"`
	UpdatedAt     int    `form:"updated_at"`
}
type CartsUpdateResp struct {
	Result Result `json:"result"`
}
type CartsDeleteResp struct {
	Result Result `json:"result"`
}
