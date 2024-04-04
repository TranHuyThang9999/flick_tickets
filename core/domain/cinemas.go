package domain

import (
	"context"

	"gorm.io/gorm"
)

type Cinemas struct {
	Id         int64  `json:"id"`
	CinemaName string `json:"cinema_name"`
}
type RepositoryCinemas interface {
	AddCinema(ctx context.Context, tx *gorm.DB, req *Cinemas) error
}
