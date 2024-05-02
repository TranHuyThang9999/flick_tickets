package domain

import (
	"context"

	"gorm.io/gorm"
)

type ShowTime struct {
	ID             int64   `json:"id"`
	TicketID       int64   `json:"ticket_id"`
	CinemaName     string  `json:"cinema_name"`
	MovieTime      int     `json:"movie_time"` //string
	SelectedSeat   string  `json:"selected_seat"`
	Quantity       int     `json:"quantity"`
	OriginalNumber int     `json:"original_number"`
	CreatedAt      int     `json:"created_at"`
	UpdatedAt      int     `json:"updated_at"`
	Price          float64 `json:"price"`
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
	GetInformationShowTimeForTicketByTicketId(ctx context.Context, id int64) (*ShowTime, error)
	UpdateQuantitySeat(ctx context.Context, tx *gorm.DB, showTimeId int64, quantity int, selected_seat string) error
	GetShowTimeByNameCinema(ctx context.Context, cinema_name string) ([]*ShowTime, error)
	GetListShowTimeByListId(ctx context.Context, ids []int64) ([]*ShowTime, error)
}
