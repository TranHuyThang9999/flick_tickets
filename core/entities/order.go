package entities

import "flick_tickets/core/domain"

type OrdersReq struct {
	Id         int64  `form:"id"`
	ShowTimeId int64  `form:"show_time_id"`
	Email      string `form:"email"`
	Seats      string `form:"seats"`
}
type OrdersResponseResgister struct {
	Result  Result `json:"result"`
	OrderId int64  `json:"orderId"`
	Created int64  `json:"created"`
}
type OrdersResponseGetById struct {
	Result  Result         `json:"result"`
	Orders  *domain.Orders `json:"orders"`
	Created int64          `json:"created"`
}
type OrderSendTicketToEmail struct {
	ID         int64   `form:"id"`
	MoviceName string  `form:"movice_name"`
	MovieTime  int     `form:"movie_time"`
	Price      float64 `form:"price"`
	Seats      string  `form:"seats"`
	CinemaName string  `form:"cinema_name"`
}
type OrderReqUpSert struct {
	Id    int64  `form:"id"`
	Email string `form:"email"`
}
type OrderRespUpSert struct {
	Result  Result `json:"result"`
	Created int64  `json:"created"`
}
type OrderSendTicketAfterPaymentReq struct {
	OrderId int64 `form:"order_id"`
}

type OrderSendTicketAfterPaymentResp struct {
	Result  Result `json:"result"`
	Created int64  `json:"created"`
}
