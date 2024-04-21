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
	CreatedAt     int     `json:"created_at"`
	UpdatedAt     int     `json:"updated_at"`
}
type TicketreqFindByForm struct {
	ID            int64  `form:"id"`
	Name          string `form:"name"`
	Description   string `form:"description"`
	Sale          int    `form:"sale"`
	ReleaseDate   int    `form:"release_date"`
	Status        int    `form:"status"`
	MovieDuration int    `form:"movieDuration"`
	AgeLimit      int    `form:"age_limit"`
	Director      string `form:"director"`
	Actor         string `form:"actor"`
	Producer      string `form:"producer"`
	CreatedAt     int    `form:"created_at"`
	UpdatedAt     int    `form:"updated_at"`
}

type TicketReqUpdateById struct {
	ID            int64   `json:"id"`
	Name          string  `json:"name"`
	Price         float64 `json:"price"`
	MaxTicket     int64   `json:"max_ticket"`
	Quantity      int     `json:"quantity"`
	Description   string  `json:"description"`
	Sale          int     `json:"sale"`
	Showtime      string  `json:"showtime"`
	MovieDuration int     `json:"movieDuration"`
	AgeLimit      int     `json:"age_limit"`
	ReleaseDate   int     `json:"release_date"`
	Director      string  `json:"director"`
	Actor         string  `json:"actor"`
	Producer      string  `json:"producer"`
	CreatedAt     int     `json:"created_at"`
	UpdatedAt     int     `json:"updated_at"`
}
type RepositoryTickets interface {
	AddTicket(ctx context.Context, tx *gorm.DB, req *Tickets) error
	GetAllTickets(ctx context.Context, req *TicketreqFindByForm) ([]*Tickets, error)
	UpdateTicketById(ctx context.Context, tx *gorm.DB, req *TicketReqUpdateById) error
	DeleteTicketsById(ctx context.Context, tx *gorm.DB, id int64) error
	UpdateTicketQuantity(ctx context.Context, tx *gorm.DB, id int64, quantity int) error
	GetTicketById(ctx context.Context, id int64) (*Tickets, error)
	// UpdateTicketSelectedSeat(ctx context.Context, tx *gorm.DB, id int64, selected_seat string) error
}
