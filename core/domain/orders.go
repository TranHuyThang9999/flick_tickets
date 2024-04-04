package domain

import (
	"context"

	"gorm.io/gorm"
)

type Orders struct {
	ID          int64   `json:"id"`
	TicketID    int64   `json:"ticket_id"`
	Showtime    string  `json:"showtime"`
	ReleaseDate int     `json:"release_date"`
	Email       string  `json:"email"`
	OTP         int64   `json:"otp"`
	Description string  `json:"description"`
	Status      int     `json:"status"`
	Price       float64 `json:"price"`
	Seats       int     `json:"seats"`
	Sale        int     `json:"sale"`
	CreatedAt   int     `json:"created_at"`
	UpdatedAt   int     `json:"updated_at"`
}
type RepositoryOrder interface {
	RegisterTicket(ctx context.Context, tx *gorm.DB, req *Orders) error
	CancelTicket(ctx context.Context, id int64) error // update
	GetOrderById(ctx context.Context, id int64) (*Orders, error)
}
