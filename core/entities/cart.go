package entities

import "flick_tickets/core/domain"

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
	Result Result          `json:"result"`
	Carts  []*domain.Carts `json:"carts"`
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
