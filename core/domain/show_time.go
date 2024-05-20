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
	Discount       int     `json:"discount"`
}
type ShowTimeDelete struct {
	ID        int64 `json:"id"`
	TicketID  int64 `json:"ticket_id"`
	MovieTime int   `json:"movie_time"`
}
type ShowTimeCheckList struct {
	//TicketID   int64  `form:"ticket_id"`
	CinemaName string `json:"cinema_name"`
	MovieTime  int    `json:"movie_time"`
}
type ShowTimeUpdateReq struct {
	ID             int64   `json:"id"`
	TicketID       int64   `json:"ticket_id"`
	CinemaName     string  `json:"cinema_name"`
	MovieTime      int     `json:"movie_time"` //string
	Quantity       int     `json:"quantity"`
	OriginalNumber int     `json:"original_number"`
	Price          float64 `json:"price"`
	UpdatedAt      int     `json:"updated_at"`
	Discount       int     `json:"discount"`
}
type RepositoryShowTime interface {
	AddShowTime(ctx context.Context, req *ShowTime) error
	AddListShowTime(ctx context.Context, tx *gorm.DB, req []*ShowTime) error
	DeleteShowTimeByTicketId(ctx context.Context, req *ShowTimeDelete) error                 // ko dung
	GetTimeUseCheckAddTicket(ctx context.Context, req *ShowTimeCheckList) (*ShowTime, error) //.ko dung
	FindDuplicateShowTimes(ctx context.Context, movieTimes []int, cinemaName []string) ([]*ShowTime, error)
	GetShowTimeByTicketId(ctx context.Context, ticketId int64) ([]*ShowTime, error)
	GetAll(ctx context.Context) ([]*ShowTime, error)
	GetInformationShowTimeForTicketByTicketId(ctx context.Context, id int64) (*ShowTime, error)
	UpdateQuantitySeat(ctx context.Context, tx *gorm.DB, showTimeId int64, quantity int, selected_seat string) error
	GetShowTimeByNameCinema(ctx context.Context, cinema_name string) ([]*ShowTime, error)
	GetListShowTimeByListId(ctx context.Context, ids []int64) ([]*ShowTime, error)
	DeleteShowTimesByTicketId(ctx context.Context, tx *gorm.DB, ticketId int64) error
	UpsertListShowTime(ctx context.Context, req []*ShowTime) error
	DeleteShowTimeByid(ctx context.Context, tx *gorm.DB, show_time_id int64) error
	GetShowTimeById(ctx context.Context, show_time_id int64) (*ShowTime, error)
	UpdateShowTimeById(ctx context.Context, req *ShowTimeUpdateReq) error
	FindDuplicateShowTimeUseUpdate(ctx context.Context, movieTime int, cinemaName string) ([]*ShowTime, error)
	GetAllShowTime(ctx context.Context) ([]*ShowTime, error)
}
