package entities

import "flick_tickets/core/domain"

type OrdersReq struct {
	Id         int64  `form:"id"`
	ShowTimeId int64  `form:"show_time_id"`
	Email      string `form:"email"`
	Seats      string `form:"seats"` //list [2,4,5]
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
type OrderCancelBtyIdreq struct {
	OrderId int64 `form:"order_id"`
}
type OrderCancelBtyIdresp struct {
	Result Result `json:"result"`
}
type OrderGetAll struct {
	Result Result           `json:"result"`
	Total  int              `json:"total"`
	Orders []*domain.Orders `json:"orders"`
}

type OrderHistoryReq struct {
	Email string `form:"email"`
}
type OrderHistoryEntities struct {
	ID             int64   `json:"id"`
	MovieName      string  `json:"movie_name"` // ten phim
	CinemaName     string  `json:"cinema_name"`
	Email          string  `json:"email"`
	ReleaseDate    int     `json:"release_date"`
	Description    string  `json:"description"`
	Status         int     `json:"status"`
	Price          float64 `json:"price"`
	Seats          string  `json:"seats"`
	MovieTime      int     `json:"movie_time"`
	AddressDetails string  `json:"address_details"`
	CreatedAt      int     `json:"created_at"`
}
type OrderHistoryResp struct {
	Result               Result                  `json:"result"`
	OrderHistoryEntities []*OrderHistoryEntities `json:"order_history_entities"`
}

type OrderRevenueReq struct {
	CinemaName        string `form:"cinema_name"`
	MovieName         string `form:"movie_name"`
	TimeDistanceStart int    `form:"time_distance_start"`
	TimeDistanceEnd   int    `form:"time_distance_end"`
}
type OrderRevenueResp struct {
	Result Result  `json:"result"`
	Sum    float64 `json:"sum"`
}
type OrderGetAllFromOrderResp struct {
	Result Result           `json:"result"`
	Orders []*domain.Orders `json:"orders"`
}
type OrderGetAllFromOrderByCinemaNameResp struct {
	Result Result           `json:"result"`
	Orders []*domain.Orders `json:"orders"`
}
type OrderStatisticalReq struct {
	StartTime int `form:"start_time"`
	EndTime   int `form:"end_time"`
}
type OrderStatisticalResp struct {
	Result Result           `json:"result"`
	Orders []*domain.Orders `json:"orders"`
}
