package domain

import (
	"context"

	"gorm.io/gorm"
)

// Struct cho bảng tickets
type Tickets struct {
	ID          int64   `json:"id"`
	UserId      int64   `json:"user_id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
	Description string  `form:"description"`
	Sale        int     `json:"sale"`
	Showtime    int     `json:"showtime"`
	ReleaseDate int     `json:"release_date"`
	CreatedAt   int     `json:"created_at"`
	UpdatedAt   int     `json:"updated_at"`
}
type TicketreqFindByForm struct {
	ID          int64   `form:"id"`
	Name        string  `form:"name"`
	Price       float64 `form:"price"`
	Quantity    int     `form:"quantity"`
	Description string  `form:"description"`
	Sale        int     `form:"sale"`
	Showtime    int     `form:"showtime"`
	ReleaseDate int     `form:"release_date"`
	CreatedAt   int     `form:"created_at"`
	UpdatedAt   int     `form:"updated_at"`
}

type TicketReqUpdateById struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
	Description string  `json:"description"`
	Sale        int     `json:"sale"`
	Showtime    int     `json:"showtime"`
	ReleaseDate int     `json:"release_date"`
	CreatedAt   int     `json:"created_at"`
	UpdatedAt   int     `json:"updated_at"`
}
type RepositoryTickets interface {
	AddTicket(ctx context.Context, tx *gorm.DB, req *Tickets) error
	GetAllTickets(ctx context.Context, req *TicketreqFindByForm) ([]*Tickets, error)
	UpdateTicketById(ctx context.Context, tx *gorm.DB, req *TicketReqUpdateById) error
	DeleteTicketsById(ctx context.Context, tx *gorm.DB, id int64) error
}
