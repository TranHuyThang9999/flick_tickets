package entities

import (
	"flick_tickets/core/domain"
	"mime/multipart"
)

type Users struct {
	UserName    string                `form:"user_name"`
	Password    string                `form:"password"`
	Address     string                `form:"address"`
	Age         int                   `form:"age"`
	File        *multipart.FileHeader `form:"file"`
	Role        int                   `form:"role"`
	ExpiredTime int                   `form:"expired_time"`
}
type UserResp struct {
	Result  Result `json:"result"`
	Created int    `json:"created"`
}
type UserRequestFindByForm struct {
	Id        int64  `form:"id"`
	UserName  string `form:"user_name"`
	Age       int    `form:"age"`
	Address   string `form:"address"`
	CreatedAt int    `form:"created_at"`
	UpdatedAt int    `form:"updated_at"`
}
type UserRespFindByForm struct {
	Users   []*domain.Users `json:"users"`
	Result  Result          `json:"result"`
	Created int             `json:"created"`
}
type LoginReq struct {
	UserName string `form:"user_name" validate:"required"`
	Password string `form:"password" validate:"required"`
}

type ShowTimeUpdateByIdReq struct {
	ID         int64   `form:"id"`
	TicketID   int64   `form:"ticket_id"`
	CinemaName string  `form:"cinema_name"`
	MovieTime  int     `form:"movie_time"` //string
	Quantity   int     `form:"quantity"`
	Price      float64 `form:"price"`
	Discount   int     `form:"discount"`
}
