package entities

import "flick_tickets/core/domain"

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
	Result    Result             `json:"result"`
	Showtimes []*domain.ShowTime `json:"showtimes"`
}
