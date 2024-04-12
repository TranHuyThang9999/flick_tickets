package entities

import "flick_tickets/core/domain"

type OrdersReq struct {
	TicketId   int64  `form:"ticket_id"`
	Email      string `form:"email"`
	Seats      int    `form:"seats"`
	CinemaName string `form:"cinema_name"`
	MovieTime  int    `form:"movie_time"`
}
type OrdersResponseResgister struct {
	Result  Result `json:"result"`
	Created int64  `json:"created"`
}
type OrdersResponseGetById struct {
	Result  Result         `json:"result"`
	Orders  *domain.Orders `json:"orders"`
	Created int64          `json:"created"`
}
type OrderSendTicketToEmail struct {
	ID          int64   `form:"id"`
	MoviceName  string  `form:"moviceName"`
	ReleaseDate int     `form:"release_date"`
	Price       float64 `form:"price"`
	Seats       int     `form:"seats"`
	CinemaName  string  `form:"cinemaName"`
}
