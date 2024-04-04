package domain

import "context"

type ShowTime struct {
	ID         int64  `json:"id"`
	TicketID   int64  `json:"ticket_id"`
	CinemaName string `json:"cinema_name"`
	MovieTime  string `json:"movie_time"` //string
	CreatedAt  int    `json:"created_at"`
	UpdatedAt  int    `json:"updated_at"`
}
type ShowTimeDelete struct {
	ID        int64 `json:"id"`
	TicketID  int64 `json:"ticket_id"`
	MovieTime int   `json:"movie_time"`
}
type ShowTimeCheckList struct {
	CinemaName string `json:"cinema_name"`
	MovieTime  string `json:"movie_time"`
}
type RepositoryShowTime interface {
	AddShowTime(ctx context.Context, req *ShowTime) error
	DeleteShowTimeByTicketId(ctx context.Context, req *ShowTimeDelete) error
	GetTimeUseCheckAddTicket(ctx context.Context, req *ShowTimeCheckList) ([]*ShowTime, error)
}
