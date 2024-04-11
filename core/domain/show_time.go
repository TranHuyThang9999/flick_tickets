package domain

import (
	"context"

	"gorm.io/gorm"
)

type ShowTime struct {
	ID         int64  `json:"id"`
	TicketID   int64  `json:"ticket_id"`
	CinemaName string `json:"cinema_name"`
	MovieTime  int    `json:"movie_time"` //string
	CreatedAt  int    `json:"created_at"`
	UpdatedAt  int    `json:"updated_at"`
}
type ShowTimeDelete struct {
	ID        int64 `json:"id"`
	TicketID  int64 `json:"ticket_id"`
	MovieTime int   `json:"movie_time"`
}
type ShowTimeCheckList struct {
	//TicketID   int64  `form:"ticket_id"`
	CinemaName string `form:"cinema_name"`
	MovieTime  int    `form:"movie_time"`
}
type RepositoryShowTime interface {
	AddShowTime(ctx context.Context, req *ShowTime) error
	AddListShowTime(ctx context.Context, tx *gorm.DB, req []*ShowTime) error
	DeleteShowTimeByTicketId(ctx context.Context, req *ShowTimeDelete) error
	GetTimeUseCheckAddTicket(ctx context.Context, req *ShowTimeCheckList) ([]*ShowTime, error) //.ko dung

	FindDuplicateShowTimes(ctx context.Context, movieTimes []int, cinemaName []string) ([]*ShowTime, error)
	GetShowTimeByTicketId(ctx context.Context, ticketId int64) ([]*ShowTime, error)
	GetAll(ctx context.Context) ([]*ShowTime, error)
}
