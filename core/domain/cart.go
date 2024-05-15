package domain

import "context"

type Carts struct {
	Id            int64   `json:"id"`
	UserName      string  `json:"user_name"`
	ShowTimeId    int64   `json:"show_time_id"`
	SeatsPosition string  `json:"seats_position"`
	Price         float64 `json:"price"`
	CreatedAt     int     `json:"created_at"`
	UpdatedAt     int     `json:"updated_at"`
}
type CartFindByFormReq struct {
	Id            int64   `form:"id"`
	UserName      string  `form:"user_name"`
	ShowTimeId    int64   `form:"show_time_id"`
	SeatsPosition string  `form:"seats_position"`
	Price         float64 `form:"price"`
	// CreatedAt     int     `form:"created_at"`
	// UpdatedAt     int     `form:"updated_at"`
}
type RepositoryCarts interface {
	AddCarts(ctx context.Context, req *Carts) error
	UpdateCartById(ctx context.Context, req *Carts) error
	DeleteCartById(ctx context.Context, cartId int64) error
	FindCartByForm(ctx context.Context, req *CartFindByFormReq) ([]*Carts, error)
	DeleteCartByShowTimeId(ctx context.Context, show_time_id int64) error
}
