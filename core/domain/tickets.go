package domain

import (
	"context"

	"gorm.io/gorm"
)

// Struct cho bảng tickets
type Tickets struct {
	ID            int64   `json:"id"`
	Name          string  `json:"name"`
	Price         float64 `json:"price"`
	Description   string  `json:"description"`
	Sale          int     `json:"sale"`
	ReleaseDate   int     `json:"release_date"`
	Status        int     `json:"status"`
	MovieDuration int     `json:"movieDuration"`
	AgeLimit      int     `json:"age_limit"`
	Director      string  `json:"director"`
	Actor         string  `json:"actor"`
	Producer      string  `json:"producer"`
	MovieType     string  `json:"movie_type"`
	CreatedAt     int     `json:"created_at"`
	UpdatedAt     int     `json:"updated_at"`
}
type TicketreqFindByForm struct {
	ID               int64   `form:"id"`
	Name             string  `form:"name"`
	Price            float64 `form:"price"`
	Description      string  `form:"description"`
	Sale             int     `form:"sale"`
	ReleaseDate      int     `form:"release_date"`
	Status           int     `form:"status"`
	MovieDuration    int     `form:"movieDuration"`
	AgeLimit         int     `form:"age_limit"`
	Director         string  `form:"director"`
	Actor            string  `form:"actor"`
	Producer         string  `form:"producer"`
	MovieType        string  `form:"movie_type"`
	CreatedAt        int     `form:"created_at"`
	UpdatedAt        int     `form:"updated_at"`
	MovieTheaterName string  `form:"movie_theater_name"`
}

type TicketReqUpdateById struct {
	ID            int64   `json:"id"`
	Name          string  `json:"name"`
	Price         float64 `json:"price"`
	Description   string  `json:"description"`
	Sale          int     `json:"sale"`
	ReleaseDate   int     `json:"release_date"`
	Status        int     `json:"status"`
	MovieDuration int     `json:"movieDuration"`
	AgeLimit      int     `json:"age_limit"`
	Director      string  `json:"director"`
	Actor         string  `json:"actor"`
	Producer      string  `json:"producer"`
	MovieType     string  `json:"movie_type"`
	UpdatedAt     int     `json:"updated_at"`
}
type RepositoryTickets interface {
	AddTicket(ctx context.Context, tx *gorm.DB, req *Tickets) error
	GetAllTickets(ctx context.Context, req *TicketreqFindByForm) ([]*Tickets, error)
	UpdateTicketById(ctx context.Context, req *TicketReqUpdateById) error
	DeleteTicketsById(ctx context.Context, tx *gorm.DB, id int64) error
	UpdateTicketQuantity(ctx context.Context, tx *gorm.DB, id int64, quantity int) error
	GetTicketById(ctx context.Context, id int64) (*Tickets, error)
	// UpdateTicketSelectedSeat(ctx context.Context, tx *gorm.DB, id int64, selected_seat string) error
	GetListTicketWithSatus(ctx context.Context, staus_sale int) ([]*Tickets, error)
	GetlistTicketByListTicketId(ctx context.Context, listTicketId []int64) ([]*Tickets, error)
	GetAllTicket(ctx context.Context) ([]*Tickets, error)
	GetAllTicketsByFilmName(ctx context.Context, name string) ([]*Tickets, error)
}
