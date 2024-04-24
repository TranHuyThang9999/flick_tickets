package domain

import (
	"context"

	"gorm.io/gorm"
)

type Orders struct {
	ID             int64   `json:"id"`
	ShowTimeID     int64   `json:"show_time_id"`
	ReleaseDate    int     `json:"release_date"`
	Email          string  `json:"email"`
	OTP            int64   `json:"otp"`
	Description    string  `json:"description"`
	Status         int     `json:"status"`
	Price          float64 `json:"price"`
	Seats          string  `json:"seats"`
	Sale           int     `json:"sale"`
	MovieTime      int     `json:"movie_time"`
	AddressDetails string  `json:"addressDetails"`
	CreatedAt      int     `json:"created_at"`
	UpdatedAt      int     `json:"updated_at"`
}
type RepositoryOrder interface {
	RegisterTicket(ctx context.Context, tx *gorm.DB, req *Orders) error
	CancelTicket(ctx context.Context, id int64) error // update
	GetOrderById(ctx context.Context, id int64) (*Orders, error)
	UpsertOrder(ctx context.Context, email string, orderId int64, status int) error
	GetOrderByEmail(ctx context.Context, email string) (*Orders, error)
}
