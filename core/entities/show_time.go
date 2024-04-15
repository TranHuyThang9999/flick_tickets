package entities

type ShowTimeAddReq struct {
	TicketID   int64  `form:"ticket_id"`
	CinemaName string `form:"cinema_name"`
	MovieTime  int    `form:"movie_time"`
}
type ShowTimeDelete struct {
	ID        int64 `form:"id"`
	TicketID  int64 `form:"ticket_id"`
	MovieTime int   `form:"movie_time"`
}
type ShowTimeAddResponse struct {
	Result    Result `json:"result"`
	CreatedAt int    `json:"created_at"`
}
type ShowTimeDeleteResponse struct {
	Result    Result `json:"result"`
	CreatedAt int    `json:"created_at"`
}
type ShowTimeByTicketIdresp struct {
	Result    Result      `json:"result"`
	Showtimes []*ShowTime `json:"showtimes"`
}

type ShowTime struct {
	ID              int64  `json:"id"`
	TicketID        int64  `json:"ticket_id"`
	CinemaName      string `json:"cinema_name"`
	MovieTime       int    `json:"movie_time"` //string
	Description     string `json:"description"`
	Conscious       string `json:"conscious"`
	District        string `json:"district"`
	Commune         string `json:"commune"`
	AddressDetails  string `json:"address_details"`
	WidthContainer  int    `json:"width_container"`
	HeightContainer int    `json:"height_container"`
}
